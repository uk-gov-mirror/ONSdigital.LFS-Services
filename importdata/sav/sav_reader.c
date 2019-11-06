#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#include "sav_reader.h"

const int ACCOC_SIZE = 256 * 1024 * 1024;
const int STRUCT_ACCOC_SIZE = 1024;
const int SAV_BUFFER_SIZE = 128;  // initial buffer size for a value, will grow if necessary

struct Rows * add_row(struct Data *data)  {
    printf("ADDING ROW\n");
    data->rows[data->row_count] = malloc(sizeof(struct Rows));
    data->rows[data->row_count]->row_length = 0;
    data->row_count++;
    data->row_position = 0;

    return data->rows[data->row_count - 1];
}

void add_to_row(struct Data *data, const char *value) {
    printf("ADDING TO ROW\n");
    struct Rows *current_row = data->rows[data->row_count - 1];
    int position = data->row_position;

    current_row->row_data[position] = malloc(sizeof(value + 1));
    strcpy(current_row->row_data[position], value);
    data->row_position++;
    current_row->row_length++;
}

//int handle_metadata(readstat_metadata_t *metadata, void *ctx) {
//    struct Data *data = (struct Data *) ctx;
//    data->var_count = readstat_get_var_count(metadata);
//    return READSTAT_HANDLER_OK;
//}

int handle_variable(int index, readstat_variable_t *variable, const char *val_labels, void *ctx) {
    struct Data *data = (struct Data *) ctx;
    const char *var_name = readstat_variable_get_name(variable);
    const char *var_description = readstat_variable_get_label(variable);

    int header_size = sizeof(struct Header);
    data->header = realloc(data->header, (data->header_count * header_size)  + header_size);

    struct Header *header = malloc(sizeof(struct Header));
    data->header[data->header_count] = header;

    unsigned long var_len = strlen(var_name) + 1;
    header->var_name = malloc(var_len);

    strcpy(header->var_name, var_name);

    if (var_description != NULL) {
        unsigned long desc_len = strlen(var_description) + 1;
        header->var_description = malloc(desc_len);
        strcpy(header->var_description, var_description);
    }

    header->var_type = readstat_variable_get_type(variable);
    header->length = readstat_variable_get_storage_width(variable);
    header->length = variable->decimals;

    data->header_count++;

    return READSTAT_HANDLER_OK;
}

int handle_value(int obs_index, readstat_variable_t *variable, readstat_value_t value, void *ctx) {

    return READSTAT_HANDLER_OK;

    struct Data *data = (struct Data *) ctx;
    int var_index = readstat_variable_get_index(variable);

    struct Rows *current_row;

    if (var_index == 0) {
        current_row = add_row(data);
    } else {
        current_row = data->rows[data->row_count - 1];
    }

    readstat_type_t type = readstat_value_type(value);

    char *buf = data->buffer;

    switch (type) {
        case READSTAT_TYPE_STRING:

            // This will be the only place we can expect a value larger than the
            // existing SAV_BUFFER_SIZE
            // We use snprintf as it's much faster
            if (data->buffer_size <= strlen(readstat_string_value(value)) + 1) {
                data->buffer_size = strlen(readstat_string_value(value)) + SAV_BUFFER_SIZE + 1;
                data->buffer = realloc(data->buffer, data->buffer_size);
                buf = data->buffer;
            }
            char *str = (char *) readstat_string_value(value);
            for (char* p = str; (p = strchr(p, ',')) ; ++p) {
                *p = ' ';
            }
            snprintf(buf, data->buffer_size, "%s", readstat_string_value(value));
            add_to_row(data, buf);
            break;

        case READSTAT_TYPE_INT8:
            if (readstat_value_is_system_missing(value)) {
                snprintf(buf, data->buffer_size, "NaN");
            } else {
                snprintf(buf, data->buffer_size, "%d", readstat_int8_value(value));
            }
            add_to_row(data, buf);
            break;

        case READSTAT_TYPE_INT16:
            if (readstat_value_is_system_missing(value)) {
                snprintf(buf, data->buffer_size, "NaN");
            } else {
                snprintf(buf, data->buffer_size, "%d", readstat_int16_value(value));
            }
            add_to_row(data, buf);
            break;

        case READSTAT_TYPE_INT32:
            if (readstat_value_is_system_missing(value)) {
                snprintf(buf, data->buffer_size, "Nan");
            } else {
                snprintf(buf, data->buffer_size, "%d", readstat_int32_value(value));
            }
            add_to_row(data, buf);
            break;

        case READSTAT_TYPE_FLOAT:
            if (readstat_value_is_system_missing(value)) {
                snprintf(buf, data->buffer_size, "NaN");
            } else {
                snprintf(buf, data->buffer_size, "%f", readstat_float_value(value));
            }
            add_to_row(data, buf);

            break;

        case READSTAT_TYPE_DOUBLE:
            if (readstat_value_is_system_missing(value)) {
                snprintf(buf, data->buffer_size, "NaN");
            } else {
                snprintf(buf, data->buffer_size, "%lf", readstat_double_value(value));
            }
            add_to_row(data, buf);

            break;

        default:
            return READSTAT_HANDLER_OK;
    }

    return READSTAT_HANDLER_OK;
}

void cleanup(struct Data *data) {
   for (int i = 0; i < data->header_count; i++) {
        struct Header *header = data->header[i];
        if (header->var_name != NULL) free(header->var_name);
        if (header->var_description != NULL) free(header->var_description);
        free(header);
   }

   for (int i = 0; i < data->row_count; i++) {
       struct Rows *rows = data->rows[i];
       for (int j = 0; j < rows->row_length; j++) {
          if (rows->row_data[j] != 0) free(rows->row_data[j]);
       }
       free(rows);
   }
}

struct Data * parse_sav(const char *input_file) {

    if (input_file == 0) {
        return NULL;
    }

    readstat_error_t error;
    readstat_parser_t *parser = readstat_parser_init();
//    readstat_set_metadata_handler(parser, &handle_metadata);
    readstat_set_variable_handler(parser, &handle_variable);
    readstat_set_value_handler(parser, &handle_value);

    struct Data *sav_data = (struct Data *) malloc(sizeof(struct Data));
    sav_data->rows = NULL;
    sav_data->row_count = 0;
    sav_data->row_position = 0;

    sav_data->buffer = malloc(SAV_BUFFER_SIZE);
    sav_data->buffer_size = SAV_BUFFER_SIZE;

    sav_data->header = NULL;
    sav_data->header_count = 0;

    error = readstat_parse_sav(parser, input_file, sav_data);

    readstat_parser_free(parser);

    if (error != READSTAT_OK) {
      return NULL;
    }

    return sav_data;

}

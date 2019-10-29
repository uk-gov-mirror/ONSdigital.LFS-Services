-- MySQL Script generated by MySQL Workbench
-- Fri Oct 25 13:34:09 2019
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0;
SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE =
        'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema lfs
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `lfs`;

-- -----------------------------------------------------
-- Schema lfs
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `lfs` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `lfs`;

create table addresses
(
    id                  int unsigned auto_increment,
    pcd7                varchar(7)  not null,
    tlec99              varchar(3)  null,
    ELWA                decimal(38) null,
    SCOTER              varchar(6)  null,
    Walespca            decimal(38) null,
    ward03              varchar(6)  null,
    scotpca             decimal(38) null,
    ukpca               decimal(38) null,
    TTWA07              decimal(38) null,
    ttwa08              decimal(38) null,
    pca2010             varchar(3)  null,
    nuts2               varchar(4)  null,
    nuts3               varchar(5)  null,
    nuts4               varchar(7)  null,
    nuts10              varchar(10) null,
    nuts102             varchar(4)  null,
    nuts103             varchar(5)  null,
    nuts104             varchar(7)  null,
    eregn10             varchar(2)  null,
    eregn103            varchar(3)  null,
    NUTS133             varchar(5)  null,
    NUTS132             varchar(4)  null,
    eregn133            varchar(3)  null,
    eregn13             varchar(2)  null,
    DEGURBA             decimal(38) null,
    dzone1              varchar(9)  null,
    dzone2              varchar(9)  null,
    soa1                varchar(9)  null,
    soa2                varchar(9)  null,
    ward05              varchar(6)  null,
    oacode              varchar(10) null,
    urind               decimal(38) null,
    urindsul            decimal(38) null,
    lea                 varchar(3)  null,
    ward98              varchar(6)  null,
    OSLAUA9d            varchar(9)  null,
    ctry9d              varchar(9)  not null,
    casward             varchar(6)  null,
    oa11                varchar(9)  null,
    CTY                 varchar(9)  null,
    LAUA                varchar(9)  null,
    WARD                varchar(9)  null,
    CED                 varchar(9)  null,
    GOR9d               varchar(9)  null,
    PCON9d              varchar(9)  null,
    TECLEC9d            varchar(9)  null,
    TTWA9d              varchar(9)  null,
    lau2                varchar(9)  null,
    PARK                varchar(9)  null,
    LSOA11              varchar(9)  null,
    MSOA11              varchar(9)  null,
    CCG                 varchar(9)  null,
    RU11IND             varchar(2)  null,
    OAC11               varchar(3)  null,
    LEP1                varchar(9)  null,
    LEP2                varchar(9)  null,
    IMD                 decimal(38) null,
    ru11indsul          decimal(38) null,
    NUTS163             varchar(5)  null,
    NUTS162             varchar(4)  null,
    eregn163            varchar(3)  null,
    eregn16             varchar(2)  not null,
    METCTY              varchar(9)  not null,
    UTLA                varchar(9)  not null,
    WIMD2014quintile    decimal(38) null,
    decile2015          decimal(38) null,
    CombinedAuthorities varchar(9)  not null,
    constraint id_UNIQUE
        unique (id)
);

alter table addresses
    add primary key (id);

create table export_definitions
(
    Variables       varchar(10) not null
        primary key,
    Research        tinyint(1)  not null,
    Regional_Client tinyint(1)  not null,
    Government      tinyint(1)  not null,
    Special_License tinyint(1)  not null,
    End_User        tinyint(1)  not null,
    Adhoc           tinyint(1)  not null
);

create table status_values
(
    id          int(11)      not null,
    description varchar(255) not null,
    constraint status_values_id_uindex
        unique (id)
);

insert into status_values(id, description)
values (0, 'Not Started');

insert into status_values(id, description)
values (1, 'File Uploaded');

insert into status_values(id, description)
values (2, 'File Reloaded');

insert into status_values(id, description)
values (3, 'Upload Failed');

alter table status_values
    add primary key (id);

create table monthly_batch
(
    id          int auto_increment,
    month       int default 0 not null,
    year        int           not null,
    status      int default 0 not null,
    description text          null,
    constraint idf_UNIQUE
        unique (id),
    constraint monthly_batch_status_values_id_fk
        foreign key (id) references status_values (id)
);

alter table monthly_batch
    add primary key (id);

create table annual_batch
(
    id          int        not null,
    year        int        null,
    status      int        null,
    description mediumtext null,
    constraint id_UNIQUE
        unique (id),
    constraint ab_to_mb
        foreign key (id) references monthly_batch (id),
    constraint annual_batch_status_values_id_fk
        foreign key (id) references status_values (id)
);

alter table annual_batch
    add primary key (id);

create table gb_batch_items
(
    id     int not null,
    year   int null,
    month  int null,
    week   int not null,
    status int null,
    primary key (week, id),
    constraint batch
        foreign key (id) references monthly_batch (id),
    constraint gb_batch_items_status_values_id_fk
        foreign key (id) references status_values (id)
);

create table ni_batch_item
(
    id     int not null,
    year   int null,
    month  int null,
    status int null,
    constraint id_UNIQUE
        unique (id),
    constraint monthly
        foreign key (id) references monthly_batch (id),
    constraint ni_batch_item_status_values_id_fk
        foreign key (id) references status_values (id)
);

alter table ni_batch_item
    add primary key (id);

create table quarterly_batch
(
    id          int auto_increment,
    quarter     int        null,
    year        int        null,
    status      int        null,
    description mediumtext null,
    constraint id_UNIQUE
        unique (id),
    constraint qb_to_mb
        foreign key (id) references monthly_batch (id),
    constraint quarterly_batch_status_values_id_fk
        foreign key (id) references status_values (id)
);

alter table quarterly_batch
    add primary key (id);

create table survey
(
    id            int          not null,
    file_name     varchar(255) not null,
    file_source   char(2)      null,
    week          int          null,
    month         int          null,
    year          int          null,
    column_name   varchar(255) not null,
    column_number int          not null,
    kind          int(255)     not null,
    column_rows   longtext     not null,
    primary key (id, file_name, column_name),
    constraint gb_key
        foreign key (id) references gb_batch_items (id)
            on delete cascade,
    constraint ni_key
        foreign key (id) references ni_batch_item (id)
            on delete cascade
);

create table survey_audit
(
    id             int                        not null,
    file_name      varchar(1024) charset utf8 not null,
    file_source    char(2)                    not null,
    week           int                        null,
    month          int                        null,
    year           int                        null,
    reference_date datetime                   not null,
    num_var_file   int                        not null default 0,
    num_var_loaded int                        not null default 0,
    num_ob_file    int                        not null default 0,
    num_ob_loaded  int                        not null default 0,
    status         int                        not null,
    message        varchar(1024)              null,
    constraint survey_audit_status_values_id_fk
        foreign key (status) references status_values (id)
);

create index survey_audit_file_name_index
    on survey_audit (file_name);

create table users
(
    username varchar(255) not null,
    password varchar(255) not null,
    constraint users_username_uindex
        unique (username)
);

alter table users
    add primary key (username);

insert into users(username, password)
values ('Admin', '$2a$04$Su7c9o6E9pLaGut2Nv9FqO2ZUbntDmUweOlO/Vj3hczi86qrnbKK2');

SET SQL_MODE = @OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;

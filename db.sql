CREATE TABLE company (
                                      id varchar(45) NOT NULL,
                                      code varchar(20) NULL,
                                      name varchar(100) NOT NULL,
                                      status int4 NOT NULL,
                                      create_by varchar(45) NOT NULL,
                                      create_date timestamptz NOT NULL,
                                      update_by varchar(45) NOT NULL,
                                      update_date timestamptz NOT NULL,
                                      CONSTRAINT company_pk PRIMARY KEY (id)
);
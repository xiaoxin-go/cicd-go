drop table if exists t_resource_template;
create table t_resource_template
(
    id          int auto_increment primary key,
    name        varchar(50) unique comment '模板名',
    cpu         float comment 'cpu数量',
    memory      int comment '内存大小',
    memory_unit varchar(20) comment '内存单位, Mi, Gi',
    description varchar(255) comment '模板描述',
    created_by  varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    created_at  datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_by  varchar(32)          DEFAULT '' COMMENT '更新人',
    updated_at  datetime             DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) comment '资源模板表' charset 'utf8';

drop table if exists t_app;
create table t_app
(
    id           int auto_increment primary key,
    name         varchar(64) unique comment '应用名',
    service_type varchar(20) comment '应用类型(stated,stateless)',
    description  varchar(255) comment '描述',
    created_by   varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    created_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_by   varchar(32)          DEFAULT '' COMMENT '更新人',
    updated_at   datetime             DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) comment '应用表' charset 'utf8';

drop table if exists t_healthcheck_template;
create table t_healthcheck_template
(
    id              int auto_increment primary key,
    name            varchar(50) unique comment '健康检查模板名',
    readiness_probe varchar(255) comment 'readiness_probe',
    liveness_probe  varchar(255) comment 'liveness_probe',
    startup_probe   varchar(255) comment 'startup_probe',
    description  varchar(255) comment '描述',
    created_by      varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    created_at      datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_by      varchar(32)          DEFAULT '' COMMENT '更新人',
    updated_at      datetime             DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) comment '健康检查模板' charset 'utf8';
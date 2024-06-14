drop table if exists t_resource_template;
create table t_resource_template
(
    id         int auto_increment primary key,
    name varchar(50) unique comment '模板名',
    cpu float comment 'cpu数量',
    memory int comment '内存大小',
    memory_unit varchar(20) comment '内存单位, Mi, Gi',
    description varchar(255) comment '模板描述',
    created_by varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    created_at datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_by varchar(32)          DEFAULT '' COMMENT '更新人',
    updated_at datetime             DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) comment '资源模板表' charset 'utf8';
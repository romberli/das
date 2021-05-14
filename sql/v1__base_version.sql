CREATE TABLE `t_meta_app_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `app_name` varchar(100) NOT NULL COMMENT '应用系统名称',
  `level` tinyint(4) NOT NULL COMMENT '系统等级: 1-A, 2-B, 3-C',
  `owner_id` int(11) DEFAULT NULL COMMENT '应用系统主要负责人ID',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_app_name` (`app_name`),
  KEY `idx02_level` (`level`),
  KEY `idx03_owner_id` (`owner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用系统信息表';

CREATE TABLE `t_meta_app_db_map` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `app_id` int(11) NOT NULL COMMENT '应用系统ID',
  `db_id` int(11) NOT NULL COMMENT '数据库ID',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_app_id_db_id` (`app_id`,`db_id`),
  KEY `idx02_db_id` (`db_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用系统-数据库映射表';

CREATE TABLE `t_meta_db_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `db_name` varchar(100) NOT NULL COMMENT '数据库名称',
  `cluster_id` int(11) NOT NULL COMMENT '数据库集群ID',
  `cluster_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '集群类型: 1-单库, 2-分库分表',
  `owner_id` int(11) DEFAULT NULL COMMENT '数据库主要负责人ID',
  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_db_name_cluster_id_cluster_type` (`db_name`,`cluster_id`,`cluster_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据库信息表';

CREATE TABLE `t_meta_env_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `env_name` varchar(100) NOT NULL COMMENT '环境名称',
  `del_flag` tinyint(4) DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_env_name` (`env_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='环境信息表';

CREATE TABLE `t_meta_middleware_cluster_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_name` varchar(100) NOT NULL COMMENT '中间件集群名称',
  `owner_id` int(11) DEFAULT NULL COMMENT '中间件主要负责人ID',
  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_cluster_name` (`cluster_name`),
  KEY `idx02_owner_id` (`owner_id`),
  KEY `idx03_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='中间件集群信息表';

CREATE TABLE `t_meta_middleware_server_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_id` int(11) NOT NULL COMMENT '中间件集群ID',
  `server_name` varchar(100) NOT NULL COMMENT '中间件服务名称',
  `middleware_role` tinyint(4) NOT NULL COMMENT '中间件角色: 1-rw, 2-ro, 3-das',
  `host_ip` varchar(100) NOT NULL COMMENT '中间件服务器IP',
  `port_num` int(11) NOT NULL COMMENT '中间件端口',
  `del_flag` tinyint(4) DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_host_ip_port_num` (`host_ip`,`port_num`),
  KEY `idx02_cluster_id_middleware_role_env_id` (`cluster_id`,`middleware_role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='中间件服务器信息表';

CREATE TABLE `t_meta_monitor_system_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `system_name` varchar(100) NOT NULL COMMENT '监控系统名称',
  `system_type` tinyint(4) NOT NULL COMMENT '监控系统类型: 1-pmm1.x, 2-pmm2.x',
  `host_ip` varchar(100) NOT NULL COMMENT '监控系统服务器IP',
  `port_num` int(11) NOT NULL COMMENT '监控系统服务器端口',
  `port_num_slow` int(11) NOT NULL COMMENT '监控系统服务器慢查询日志端口',
  `base_url` varchar(200) NOT NULL COMMENT '监控系统API入口地址',
  `env_id` tinyint(4) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_system_name` (`system_name`),
  UNIQUE KEY `idx02_host_ip_port_num` (`host_ip`,`port_num`),
  KEY `idx03_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='监控系统信息表';

CREATE TABLE `t_meta_mysql_cluster_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_name` varchar(100) NOT NULL COMMENT '集群名称',
  `middleware_cluster_id` int(11) DEFAULT NULL COMMENT '中间件集群ID',
  `monitor_system_id` int(11) DEFAULT NULL COMMENT '监控系统ID',
  `owner_id` int(11) DEFAULT NULL COMMENT '数据库集群主要负责人ID',
  `env_id` int(11) NOT NULL COMMENT '环境: 1-online, 2-rel, 3-sit, 4-uat, 5-dev',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_cluster_name` (`cluster_name`),
  KEY `idx03_monitor_system_id` (`monitor_system_id`),
  KEY `idx04_owner_id` (`owner_id`),
  key `idx05_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='MySQL集群信息表';

CREATE TABLE `t_meta_mysql_server_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cluster_id` int(11) NOT NULL COMMENT '集群ID',
  `server_name` varchar(100) NOT NULL COMMENT '数据库实例名称',
  `host_ip` varchar(100) NOT NULL COMMENT '服务器IP',
  `port_num` int(11) NOT NULL COMMENT '端口',
  `deployment_type` tinyint(4) NOT NULL COMMENT '部署方式: 1-容器, 2-物理机, 3-虚拟机',
  `version` varchar(100) DEFAULT NULL COMMENT '版本, 示例: 5.7.21',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx01_cluster_id` (`cluster_id`),
  UNIQUE KEY `idx02_host_ip_port_num` (`host_ip`,`port_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='MySQL服务器信息表';

CREATE TABLE `t_meta_user_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_name` varchar(100) NOT NULL COMMENT '姓名',
  `department_name` varchar(100) NOT NULL COMMENT '部门/团队名称',
  `employee_id` varchar(100) DEFAULT NULL COMMENT '工号',
  `account_name` varchar(100) NOT NULL COMMENT '账号名称',
  `email` varchar(100) NOT NULL COMMENT '邮箱',
  `telephone` varchar(100) DEFAULT NULL COMMENT '固定电话',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号码',
  `role` tinyint(4) NOT NULL DEFAULT '3' COMMENT '角色: 1-admin, 2-dba, 3-developer',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_employee_id` (`employee_id`),
  UNIQUE KEY `idx02_account_name` (`account_name`),
  UNIQUE KEY `idx03_email` (`email`),
  UNIQUE KEY `idx04_telephone` (`telephone`),
  UNIQUE KEY `idx05_mobile` (`mobile`),
  KEY `idx06_user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

CREATE TABLE `t_hc_default_engine_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `item_name` varchar(100) NOT NULL COMMENT '检查项名称',
  `item_weight` int NOT NULL COMMENT '权重百分比, 所有检查项项权重合计应等于100',
  `low_watermark` decimal(10, 2) NOT NULL COMMENT '低水位',
  `high_watermark` decimal(10, 2) NOT NULL COMMENT '高水位',
  `unit` decimal(10, 2) NOT NULL COMMENT '百分比, 每超过该百分比时会扣分',
  `score_deduction_per_unit_high` decimal(10, 2) NOT NULL COMMENT '高指标每单位扣分分数',
  `max_score_deduction_high` decimal(10, 2) NOT NULL COMMENT '高指标最多扣分数',
  `score_deduction_per_unit_medium` decimal(10, 2) NOT NULL COMMENT '中指标每单位扣分分数',
  `max_score_deduction_medium` decimal(10, 2) NOT NULL COMMENT '中指标最多扣分数',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_item_name` (`item_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='健康检查默认引擎配置表';

CREATE TABLE `t_hc_result` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `operation_id` int(11) NOT NULL COMMENT '操作ID',
  `weighted_average_score` int(11) NOT NULL COMMENT '加权平均分',
  `db_config_score` int(11) NOT NULL COMMENT '数据库参数配置评分',
  `db_config_data` mediumtext DEFAULT NULL COMMENT '数据库参数配置数据',
  `db_config_advice` mediumtext DEFAULT NULL COMMENT '数据库参数配置优化建议',
  `cpu_usage_score` int(11) NOT NULL COMMENT 'cpu使用率评分',
  `cpu_usage_data` mediumtext DEFAULT NULL COMMENT 'cpu使用率数据',
  `cpu_usage_high` mediumtext DEFAULT NULL COMMENT '高cpu使用率数据',
  `io_util_score` int(11) NOT NULL COMMENT 'io使用率评分',
  `io_util_data` mediumtext DEFAULT NULL COMMENT 'io使用率数据',
  `io_util_high` mediumtext DEFAULT NULL COMMENT '高io使用率数据',
  `disk_capacity_usage_score` int(11) NOT NULL COMMENT '磁盘容量使用率评分',
  `disk_capacity_usage_data` mediumtext DEFAULT NULL COMMENT '磁盘容量使用率数据',
  `disk_capacity_usage_high` mediumtext DEFAULT NULL COMMENT '高磁盘容量使用率数据',
  `connection_usage_score` int(11) NOT NULL COMMENT '连接数使用率评分',
  `connection_usage_data` mediumtext DEFAULT NULL COMMENT '连接数使用率数据',
  `connection_usage_high` mediumtext DEFAULT NULL COMMENT '高连接数使用率数据',
  `average_active_session_num_score` int(11) NOT NULL COMMENT '平均活跃会话数评分',
  `average_active_session_num_data` mediumtext DEFAULT NULL COMMENT '平均活跃会话数数据',
  `average_active_session_num_high` mediumtext DEFAULT NULL COMMENT '高平均活跃会话数数据',
  `cache_miss_ratio_score` int(11) NOT NULL COMMENT '缓存未命中率评分',
  `cache_miss_ratio_data` decimal(10, 2) DEFAULT NULL COMMENT '缓存未命中率数据',
  `cache_miss_ratio_high` decimal(10, 2) DEFAULT NULL COMMENT '高缓存未命中率数据',
  `table_size_score` int(11) NOT NULL COMMENT '表大小评分',
  `table_size_data` mediumtext DEFAULT NULL COMMENT '表大小数据',
  `table_size_high` mediumtext DEFAULT NULL COMMENT '大表数据',
  `slow_query_score` int(11) NOT NULL COMMENT '慢查询评分',
  `slow_query_data` mediumtext DEFAULT NULL COMMENT '慢查询数据',
  `slow_query_advice` mediumtext DEFAULT NULL COMMENT 'top慢查询',
  `accurate_review` tinyint(4) NOT NULL DEFAULT '0' COMMENT '准确性评价: 0-未评价, 1-准确, 2-不准确',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx01_operation_id` (`operation_id`),
  KEY `idx02_operation_id_weighted_average_score` (`operation_id`, `weighted_average_score`),
  KEY `idx03_weighted_average_score` (`weighted_average_score`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='健康检查结果表';

CREATE TABLE `t_hc_operation_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mysql_server_id` int(11) NOT NULL COMMENT 'mysql服务器ID',
  `start_time` datetime(6) NOT NULL COMMENT '检查范围开始时间',
  `end_time` datetime(6) NOT NULL COMMENT '检查范围结束时间',
  `step` int(11) NOT NULL COMMENT '采样间隔, 单位: 秒',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '运行状态: 0-未运行, 1-运行中, 2-已完成, 3-已失败',
  `message` mediumtext DEFAULT NULL COMMENT '运行日志',
  `del_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除标记: 0-未删除, 1-已删除',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `last_update_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx01_mysql_server_id_status` (`mysql_server_id`, `status`),
  KEY `idx02_start_time` (`start_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='健康检查操作表';

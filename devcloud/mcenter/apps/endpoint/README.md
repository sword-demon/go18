# 接口管理

```sql
CREATE TABLE `endpoints` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uuid` varchar(100) DEFAULT NULL,
  `service` varchar(100) DEFAULT NULL comment '服务',
  `version` varchar(100) DEFAULT NULL comment '版本',
  `resource` varchar(100) DEFAULT NULL comment '资源名称',
  `action` varchar(100) DEFAULT NULL comment '资源对于的动作即操作',
  `access_mode` tinyint(1) DEFAULT NULL comment '0-读 1-读写',
  `action_label` varchar(200) DEFAULT NULL,
  `function_name` varchar(100) DEFAULT NULL comment '对于的接口函数名称',
  `path` varchar(200) DEFAULT NULL,
  `method` varchar(100) DEFAULT NULL,
  `description` text,
  `required_auth` tinyint(1) DEFAULT NULL,
  `required_code` tinyint(1) DEFAULT NULL,
  `required_perm` tinyint(1) DEFAULT NULL,
  `required_role` json DEFAULT NULL,
  `required_audit` tinyint(1) DEFAULT NULL,
  `required_namespace` tinyint(1) DEFAULT NULL,
  `extras` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_endpoints_uuid` (`uuid`),
  KEY `idx_endpoints_created_at` (`created_at`),
  KEY `idx_endpoints_deleted_at` (`deleted_at`),
  KEY `idx_endpoints_service` (`service`),
  KEY `idx_endpoints_resource` (`resource`),
  KEY `idx_endpoints_action` (`action`),
  KEY `idx_endpoints_access_mode` (`access_mode`),
  KEY `idx_endpoints_action_label` (`action_label`),
  KEY `idx_endpoints_path` (`path`),
  KEY `idx_endpoints_method` (`method`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
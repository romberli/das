insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('db_config', 5, 0, 0, 0, 10, 50, 0, 0);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('cpu_usage', 5, 50, 70, 10, 20, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('io_usage', 5, 50, 70, 10, 20, 100, 10, 20);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('disk_capacity_usage', 20, 50, 80, 10, 40, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('connection_usage', 20, 50, 80, 10, 40, 100, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('average_active_session_num', 10, 10, 20, 5, 10, 50, 5, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('cache_miss_ratio', 5, 0.5, 2, 0.1, 20, 50, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('table_rows', 5, 10000000, 30000000, 1000000, 10, 50, 10, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('table_size', 5, 10, 30, 5, 10, 50, 10, 30);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('slow_query_execution_time', 10, 1, 10, 1, 10, 100, 5, 50);
insert into t_hc_default_engine_config(item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high, score_deduction_per_unit_medium, max_score_deduction_medium)
values('slow_query_rows_examined', 10, 100000, 1000000, 100000, 10, 100, 5, 50);

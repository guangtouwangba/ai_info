# Automating High Quality RT Planning at Scale

**URL**: http://arxiv.org/abs/2501.11803v3

## 原始摘要

Radiotherapy (RT) planning is complex, subjective, and time-intensive.
Advances with artificial intelligence (AI) promise to improve its precision and
efficiency, but progress is often limited by the scarcity of large,
standardized datasets. To address this, we introduce the Automated Iterative RT
Planning (AIRTP) system, a scalable solution for generating high-quality
treatment plans. This scalable solution is designed to generate substantial
volumes of consistently high-quality treatment plans, overcoming a key obstacle
in the advancement of AI-driven RT planning. Our AIRTP pipeline adheres to
clinical guidelines and automates essential steps, including organ-at-risk
(OAR) contouring, helper structure creation, beam setup, optimization, and plan
quality improvement, using AI integrated with RT planning software like Varian
Eclipse. Furthermore, a novel approach for determining optimization parameters
to reproduce 3D dose distributions, i.e. a method to convert dose predictions
to deliverable treatment plans constrained by machine limitations is proposed.
A comparative analysis of plan quality reveals that our automated pipeline
produces treatment plans of quality comparable to those generated manually,
which traditionally require several hours of labor per plan. Committed to
public research, the first data release of our AIRTP pipeline includes nine
cohorts covering head-and-neck and lung cancer sites to support an AAPM 2025
challenge. To our best knowledge, this dataset features more than 10 times
number of plans compared to the largest existing well-curated public dataset.
Repo: https://github.com/RiqiangGao/GDP-HMM_AAPMChallenge.


## AI 摘要

研究人员开发了自动化迭代放疗规划系统(AIRTP)，通过AI集成放疗软件实现标准化治疗计划生成。该系统自动化了危及器官勾画、辅助结构创建、束流设置等关键步骤，并提出创新方法将剂量预测转化为可执行计划。测试表明，AIRTP生成的计划质量与耗时数小时的人工计划相当。团队公开了头颈癌和肺癌的9个数据集，规模达现有最大公共数据集的10倍以上，以支持2025年AAPM挑战赛。这一突破性方案解决了AI驱动放疗规划面临的大规模标准化数据短缺问题，显著提升了放疗计划的精确性和效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T23:02:33Z
- **目录日期**: 2025-05-19

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

研究人员开发了自动化迭代放疗计划系统(AIRTP)，利用人工智能提升放疗计划的精度和效率。该系统通过自动化关键步骤(如器官轮廓勾画、射束设置和优化)，解决了传统人工计划耗时且主观的问题。AIRTP能生成与人工计划质量相当的治疗方案，同时克服了AI训练数据不足的障碍。团队发布了首个公共数据集，包含头颈癌和肺癌的9组数据，规模是现有最大公共数据集的10倍以上，将支持2025年AAPM挑战赛。该系统整合了Varian Eclipse等放疗软件，并提出将剂量预测转化为可执行计划的新方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T10:02:44Z
- **目录日期**: 2025-05-19

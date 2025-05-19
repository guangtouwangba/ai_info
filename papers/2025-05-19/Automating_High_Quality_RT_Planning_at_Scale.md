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

研究人员开发了自动化迭代放疗规划系统(AIRTP)，通过AI技术整合放疗软件，实现了从器官轮廓勾画到治疗计划优化的全流程自动化。该系统能生成大量高质量放疗方案，质量媲美人工耗时数小时制定的方案，解决了AI放疗研究中高质量数据稀缺的瓶颈问题。团队还创新性地提出了将3D剂量预测转化为可执行治疗方案的参数优化方法。作为公共研究资源，AIRTP首批发布了涵盖头颈和肺部癌症的9组数据集，规模超过现有最大公共数据集的10倍，将支持2025年AAPM挑战赛。项目代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T08:02:46Z
- **目录日期**: 2025-05-19

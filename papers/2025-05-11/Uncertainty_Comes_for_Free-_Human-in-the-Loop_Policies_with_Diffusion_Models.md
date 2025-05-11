# Uncertainty Comes for Free: Human-in-the-Loop Policies with Diffusion Models

**URL**: http://arxiv.org/abs/2503.01876v2

## 原始摘要

Human-in-the-loop (HitL) robot deployment has gained significant attention in
both academia and industry as a semi-autonomous paradigm that enables human
operators to intervene and adjust robot behaviors at deployment time, improving
success rates. However, continuous human monitoring and intervention can be
highly labor-intensive and impractical when deploying a large number of robots.
To address this limitation, we propose a method that allows diffusion policies
to actively seek human assistance only when necessary, reducing reliance on
constant human oversight. To achieve this, we leverage the generative process
of diffusion policies to compute an uncertainty-based metric based on which the
autonomous agent can decide to request operator assistance at deployment time,
without requiring any operator interaction during training. Additionally, we
show that the same method can be used for efficient data collection for
fine-tuning diffusion policies in order to improve their autonomous
performance. Experimental results from simulated and real-world environments
demonstrate that our approach enhances policy performance during deployment for
a variety of scenarios.


## AI 摘要

本文提出了一种基于扩散策略的人机协同机器人部署方法，通过计算不确定性指标让机器人仅在必要时主动寻求人工干预，减少持续人工监控的需求。该方法利用扩散策略的生成过程，在部署时动态决定是否请求操作员协助，且训练阶段无需人工交互。同时，该方法还能用于高效收集数据以微调扩散策略，提升自主性能。仿真和真实环境实验表明，该方法能有效提高多种场景下的策略性能。这一方案解决了大规模机器人部署时人工监控劳动强度高的问题，实现了更高效的人机协作。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-11T00:02:23Z
- **目录日期**: 2025-05-11

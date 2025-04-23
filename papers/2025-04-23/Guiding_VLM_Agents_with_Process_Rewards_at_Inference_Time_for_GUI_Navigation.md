# Guiding VLM Agents with Process Rewards at Inference Time for GUI Navigation

**URL**: http://arxiv.org/abs/2504.16073v1

## 原始摘要

Recent advancements in visual language models (VLMs) have notably enhanced
their capabilities in handling complex Graphical User Interface (GUI)
interaction tasks. Despite these improvements, current frameworks often
struggle to generate correct actions in challenging GUI environments.
State-of-the-art commercial VLMs are black-boxes, and fine-tuning open-source
VLMs for GUI tasks requires significant resources. Additionally, existing
trajectory-level evaluation and refinement techniques frequently fall short due
to delayed feedback and local optimization issues. To address these challenges,
we propose an approach that guides VLM agents with process supervision by a
reward model during GUI navigation and control at inference time. This guidance
allows the VLM agent to optimize actions at each inference step, thereby
improving performance in both static and dynamic environments. In particular,
our method demonstrates significant performance gains in three GUI navigation
tasks, achieving a 3.4% improvement in single step action accuracy for static
environments, along with a around 33% increase in task success rate in one
dynamic environment. With further integration of trajectory reflection and
retry mechanisms, we also demonstrate even greater enhancement in task success.


## AI 摘要

近期视觉语言模型(VLM)在图形用户界面(GUI)交互任务中取得进展，但仍面临挑战。现有商业VLM是黑箱系统，而开源模型的微调成本高昂。针对轨迹评估和优化的延迟反馈问题，研究者提出了一种新方法：在推理时使用奖励模型进行过程监督，引导VLM代理优化每一步动作。该方法在静态环境中单步动作准确率提升3.4%，在动态环境中任务成功率提高约33%。结合轨迹反思和重试机制后，任务成功率得到进一步改善。该方法有效提升了VLM在静态和动态GUI环境中的表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T13:09:00Z
- **目录日期**: 2025-04-23

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

近期视觉语言模型(VLM)在GUI交互任务中取得进展，但仍面临复杂环境下的动作生成困难。商用VLM是黑箱系统，而开源模型的微调成本高昂。现有评估方法因延迟反馈和局部优化问题效果有限。为此，研究者提出在推理时通过奖励模型进行过程监督，引导VLM代理优化每一步动作，提升静态和动态环境下的表现。该方法在三个GUI导航任务中显著提升性能：静态环境单步动作准确率提高3.4%，某动态环境任务成功率提升约33%。结合轨迹反思和重试机制后，任务成功率得到进一步改善。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T02:29:15Z
- **目录日期**: 2025-04-24

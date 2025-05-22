# HCRMP: A LLM-Hinted Contextual Reinforcement Learning Framework for Autonomous Driving

**URL**: http://arxiv.org/abs/2505.15793v1

## 原始摘要

Integrating Large Language Models (LLMs) with Reinforcement Learning (RL) can
enhance autonomous driving (AD) performance in complex scenarios. However,
current LLM-Dominated RL methods over-rely on LLM outputs, which are prone to
hallucinations.Evaluations show that state-of-the-art LLM indicates a
non-hallucination rate of only approximately 57.95% when assessed on essential
driving-related tasks. Thus, in these methods, hallucinations from the LLM can
directly jeopardize the performance of driving policies. This paper argues that
maintaining relative independence between the LLM and the RL is vital for
solving the hallucinations problem. Consequently, this paper is devoted to
propose a novel LLM-Hinted RL paradigm. The LLM is used to generate semantic
hints for state augmentation and policy optimization to assist RL agent in
motion planning, while the RL agent counteracts potential erroneous semantic
indications through policy learning to achieve excellent driving performance.
Based on this paradigm, we propose the HCRMP (LLM-Hinted Contextual
Reinforcement Learning Motion Planner) architecture, which is designed that
includes Augmented Semantic Representation Module to extend state space.
Contextual Stability Anchor Module enhances the reliability of multi-critic
weight hints by utilizing information from the knowledge base. Semantic Cache
Module is employed to seamlessly integrate LLM low-frequency guidance with RL
high-frequency control. Extensive experiments in CARLA validate HCRMP's strong
overall driving performance. HCRMP achieves a task success rate of up to 80.3%
under diverse driving conditions with different traffic densities. Under
safety-critical driving conditions, HCRMP significantly reduces the collision
rate by 11.4%, which effectively improves the driving performance in complex
scenarios.


## AI 摘要

将大型语言模型(LLM)与强化学习(RL)结合可提升自动驾驶在复杂场景中的表现，但现有方法过度依赖易产生幻觉的LLM输出(非幻觉率仅57.95%)。本文提出LLM-Hinted RL新范式：LLM生成语义提示辅助RL进行状态增强和策略优化，同时RL通过策略学习抵消潜在错误语义指示。基于此构建的HCRMP架构包含语义增强表示模块、上下文稳定锚模块(提升多评价权重可靠性)和语义缓存模块(融合低频LLM指导与高频RL控制)。CARLA实验显示HCRMP任务成功率最高达80.3%，关键安全场景碰撞率降低11.4%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T20:01:46Z
- **目录日期**: 2025-05-22

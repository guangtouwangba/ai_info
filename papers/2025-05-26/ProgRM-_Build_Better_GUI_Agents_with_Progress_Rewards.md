# ProgRM: Build Better GUI Agents with Progress Rewards

**URL**: http://arxiv.org/abs/2505.18121v1

## 原始摘要

LLM-based (Large Language Model) GUI (Graphical User Interface) agents can
potentially reshape our daily lives significantly. However, current LLM-based
GUI agents suffer from the scarcity of high-quality training data owing to the
difficulties of trajectory collection and reward annotation. Existing works
have been exploring LLMs to collect trajectories for imitation learning or to
offer reward signals for online RL training. However, the Outcome Reward Model
(ORM) used in existing works cannot provide finegrained feedback and can
over-penalize the valuable steps in finally failed trajectories. To this end,
we propose Progress Reward Model (ProgRM) to provide dense informative
intermediate rewards by predicting a task completion progress for each step in
online training. To handle the challenge of progress reward label annotation,
we further design an efficient LCS-based (Longest Common Subsequence)
self-annotation algorithm to discover the key steps in trajectories and assign
progress labels accordingly. ProgRM is evaluated with extensive experiments and
analyses. Actors trained with ProgRM outperform leading proprietary LLMs and
ORM-trained actors, illustrating the effectiveness of ProgRM. The codes for
experiments will be made publicly available upon acceptance.


## AI 摘要

基于大语言模型（LLM）的GUI代理能显著改变日常生活，但高质量训练数据的稀缺限制了其发展。现有方法依赖LLM生成轨迹进行模仿学习或提供奖励信号，但结果奖励模型（ORM）无法提供细粒度反馈，且可能过度惩罚失败轨迹中的关键步骤。为此，本文提出进度奖励模型（ProgRM），通过预测任务完成进度为在线训练提供密集中间奖励。为解决进度标签标注难题，设计了一种基于最长公共子序列（LCS）的自标注算法，自动识别关键步骤并分配进度标签。实验表明，ProgRM训练的代理性能优于主流LLM和ORM方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T10:06:14Z
- **目录日期**: 2025-05-26

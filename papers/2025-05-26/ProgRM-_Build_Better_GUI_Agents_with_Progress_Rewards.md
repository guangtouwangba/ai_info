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

本文提出了一种基于大型语言模型（LLM）的图形用户界面（GUI）智能体训练方法——Progress Reward Model（ProgRM），以解决现有Outcome Reward Model（ORM）在在线训练中无法提供细粒度反馈的问题。ProgRM通过预测任务完成进度来生成密集的中间奖励信号，并设计了一种基于最长公共子序列（LCS）的自标注算法来自动生成进度标签。实验表明，使用ProgRM训练的智能体性能优于主流商业LLM和基于ORM训练的模型，验证了该方法的有效性。相关代码将在论文录用后公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T08:03:26Z
- **目录日期**: 2025-05-26

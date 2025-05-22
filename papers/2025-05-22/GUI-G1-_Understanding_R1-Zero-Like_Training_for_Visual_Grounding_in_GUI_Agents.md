# GUI-G1: Understanding R1-Zero-Like Training for Visual Grounding in GUI Agents

**URL**: http://arxiv.org/abs/2505.15810v1

## 原始摘要

Recent Graphical User Interface (GUI) agents replicate the R1-Zero paradigm,
coupling online Reinforcement Learning (RL) with explicit chain-of-thought
reasoning prior to object grounding and thereby achieving substantial
performance gains. In this paper, we first conduct extensive analysis
experiments of three key components of that training pipeline: input design,
output evaluation, and policy update-each revealing distinct challenges arising
from blindly applying general-purpose RL without adapting to GUI grounding
tasks. Input design: Current templates encourage the model to generate
chain-of-thought reasoning, but longer chains unexpectedly lead to worse
grounding performance. Output evaluation: Reward functions based on hit signals
or box area allow models to exploit box size, leading to reward hacking and
poor localization quality. Policy update: Online RL tends to overfit easy
examples due to biases in length and sample difficulty, leading to
under-optimization on harder cases. To address these issues, we propose three
targeted solutions. First, we adopt a Fast Thinking Template that encourages
direct answer generation, reducing excessive reasoning during training. Second,
we incorporate a box size constraint into the reward function to mitigate
reward hacking. Third, we revise the RL objective by adjusting length
normalization and adding a difficulty-aware scaling factor, enabling better
optimization on hard samples. Our GUI-G1-3B, trained on 17K public samples with
Qwen2.5-VL-3B-Instruct, achieves 90.3% accuracy on ScreenSpot and 37.1% on
ScreenSpot-Pro. This surpasses all prior models of similar size and even
outperforms the larger UI-TARS-7B, establishing a new state-of-the-art in GUI
agent grounding. The project repository is available at
https://github.com/Yuqi-Zhou/GUI-G1.


## AI 摘要

这篇论文分析了GUI智能体训练中的三个关键问题：输入设计（过长的思维链反而降低定位性能）、输出评估（基于命中信号或框面积的奖励函数易被模型利用导致定位质量差）和策略更新（在线强化学习易过拟合简单样本）。作者提出三个针对性解决方案：1）采用快速思维模板减少冗余推理；2）在奖励函数中加入框尺寸约束；3）调整RL目标函数，加入难度感知因子。最终训练的GUI-G1-3B模型在ScreenSpot和ScreenSpot-Pro数据集上分别达到90.3%和37.1%准确率，超越了同类及更大规模模型，创下新SOTA。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T14:01:12Z
- **目录日期**: 2025-05-22

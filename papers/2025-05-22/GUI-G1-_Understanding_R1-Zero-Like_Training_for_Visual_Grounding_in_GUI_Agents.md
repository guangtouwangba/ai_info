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

近期研究发现，现有GUI智能体在结合强化学习（RL）和显式思维链推理时存在三个关键问题：输入设计上，过长的思维链反而降低定位性能；输出评估中，基于点击或框面积的奖励函数易被模型利用，导致奖励作弊；策略更新时，RL对简单样本过拟合，难样本优化不足。为此，研究者提出三项改进：采用快速思维模板减少冗余推理、在奖励函数中加入框尺寸约束、调整RL目标以优化难样本。基于17K样本训练的GUI-G1-3B模型在ScreenSpot上准确率达90.3%，超越同类甚至更大模型，创下新SOTA。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T09:01:10Z
- **目录日期**: 2025-05-22

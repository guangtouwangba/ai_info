# GUI-R1 : A Generalist R1-Style Vision-Language Action Model For GUI Agents

**URL**: http://arxiv.org/abs/2504.10458v1

## 原始摘要

Existing efforts in building Graphical User Interface (GUI) agents largely
rely on the training paradigm of supervised fine-tuning on Large
Vision-Language Models (LVLMs). However, this approach not only demands
extensive amounts of training data but also struggles to effectively understand
GUI screenshots and generalize to unseen interfaces. The issue significantly
limits its application in real-world scenarios, especially for high-level
tasks. Inspired by Reinforcement Fine-Tuning (RFT) in large reasoning models
(e.g., DeepSeek-R1), which efficiently enhances the problem-solving
capabilities of large language models in real-world settings, we propose \name,
the first reinforcement learning framework designed to enhance the GUI
capabilities of LVLMs in high-level real-world task scenarios, through unified
action space rule modeling. By leveraging a small amount of carefully curated
high-quality data across multiple platforms (including Windows, Linux, MacOS,
Android, and Web) and employing policy optimization algorithms such as Group
Relative Policy Optimization (GRPO) to update the model, \name achieves
superior performance using only 0.02\% of the data (3K vs. 13M) compared to
previous state-of-the-art methods like OS-Atlas across eight benchmarks
spanning three different platforms (mobile, desktop, and web). These results
demonstrate the immense potential of reinforcement learning based on unified
action space rule modeling in improving the execution capabilities of LVLMs for
real-world GUI agent tasks.


## AI 摘要

当前构建图形用户界面(GUI)代理主要依赖大型视觉语言模型(LVLMs)的监督微调，但该方法需要大量训练数据且泛化能力有限。受强化微调(RFT)启发，研究者提出首个基于强化学习的框架，通过统一动作空间规则建模来提升LVLMs在真实GUI任务中的表现。该框架仅使用0.02%数据(3K vs 13M)，在跨8个基准测试(移动/桌面/网页)中超越现有最佳方法，展现了强化学习在提升GUI代理执行能力方面的巨大潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T23:02:26Z
- **目录日期**: 2025-04-15

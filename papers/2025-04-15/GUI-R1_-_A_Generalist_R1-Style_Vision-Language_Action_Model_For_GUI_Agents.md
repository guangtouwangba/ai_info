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

现有基于大视觉语言模型(LVLM)的图形用户界面(GUI)代理主要依赖监督微调训练范式，但这种方法需要大量训练数据且泛化能力有限。受强化微调(RFT)在大语言模型中提升问题解决能力的启发，研究者提出了首个基于强化学习的GUI代理框架，通过统一动作空间规则建模来增强LVLM在真实任务场景中的能力。该框架仅使用0.02%的数据(3K vs 13M)，在跨8个基准测试(覆盖移动、桌面和网页三大平台)上超越了现有最优方法，展现了强化学习在提升GUI代理执行能力方面的巨大潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T07:01:16Z
- **目录日期**: 2025-04-15

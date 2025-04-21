# GUI-R1 : A Generalist R1-Style Vision-Language Action Model For GUI Agents

**URL**: http://arxiv.org/abs/2504.10458v3

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

现有基于大规模视觉语言模型（LVLM）的图形用户界面（GUI）智能体主要依赖监督微调训练，但这种方法需要大量数据且泛化能力有限。受强化微调（RFT）在大模型推理任务中的启发，研究者提出了首个基于强化学习的GUI智能体框架，通过统一动作空间规则建模，仅需0.02%数据（3K vs. 13M），在跨平台（Windows/Linux/MacOS/Android/Web）的高质量数据集上使用GRPO等策略优化算法，性能即超越此前最佳方法OS-Atlas。该框架在移动、桌面和网页三大平台的8个基准测试中表现优异，展现了强化学习在提升LVLM执行真实GUI任务方面的巨大潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T21:02:04Z
- **目录日期**: 2025-04-21

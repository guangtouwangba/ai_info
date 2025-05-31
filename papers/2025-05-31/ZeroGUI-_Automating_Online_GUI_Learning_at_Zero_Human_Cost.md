# ZeroGUI: Automating Online GUI Learning at Zero Human Cost

**URL**: http://arxiv.org/abs/2505.23762v1

## 原始摘要

The rapid advancement of large Vision-Language Models (VLMs) has propelled
the development of pure-vision-based GUI Agents, capable of perceiving and
operating Graphical User Interfaces (GUI) to autonomously fulfill user
instructions. However, existing approaches usually adopt an offline learning
framework, which faces two core limitations: (1) heavy reliance on high-quality
manual annotations for element grounding and action supervision, and (2)
limited adaptability to dynamic and interactive environments. To address these
limitations, we propose ZeroGUI, a scalable, online learning framework for
automating GUI Agent training at Zero human cost. Specifically, ZeroGUI
integrates (i) VLM-based automatic task generation to produce diverse training
goals from the current environment state, (ii) VLM-based automatic reward
estimation to assess task success without hand-crafted evaluation functions,
and (iii) two-stage online reinforcement learning to continuously interact with
and learn from GUI environments. Experiments on two advanced GUI Agents
(UI-TARS and Aguvis) demonstrate that ZeroGUI significantly boosts performance
across OSWorld and AndroidLab environments. The code is available at
https://github.com/OpenGVLab/ZeroGUI.


## AI 摘要

ZeroGUI提出了一种零人工成本的在线学习框架，用于训练纯视觉GUI代理。现有方法依赖大量人工标注且难以适应动态环境，而ZeroGUI通过三个创新点突破限制：(1) 基于VLM的自动任务生成，从当前环境状态产生多样化训练目标；(2) VLM驱动的自动奖励评估，无需人工设计评价函数；(3) 两阶段在线强化学习，持续与环境交互学习。在UI-TARS和Aguvis代理上的实验表明，ZeroGUI显著提升了OSWorld和AndroidLab环境中的表现。该框架实现了完全自动化的GUI代理训练，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-31T13:07:29Z
- **目录日期**: 2025-05-31

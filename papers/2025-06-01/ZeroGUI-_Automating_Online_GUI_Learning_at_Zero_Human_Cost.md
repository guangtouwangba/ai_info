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

ZeroGUI是一个创新的在线学习框架，旨在零人工成本下自动化训练纯视觉GUI代理。它通过视觉语言模型（VLM）自动生成多样化任务目标，并评估任务完成度，无需人工标注或定制评估函数。结合两阶段在线强化学习，ZeroGUI能持续适应动态GUI环境。实验表明，该框架显著提升了UI-TARS和Aguvis等GUI代理在OSWorld和AndroidLab环境中的性能。ZeroGUI解决了现有离线方法依赖人工标注和环境适应性的局限，为GUI自动化提供了高效可扩展的解决方案。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T09:01:05Z
- **目录日期**: 2025-06-01

# GUI-Reflection: Empowering Multimodal GUI Models with Self-Reflection Behavior

**URL**: http://arxiv.org/abs/2506.08012v1

## 原始摘要

Multimodal Large Language Models (MLLMs) have shown great potential in
revolutionizing Graphical User Interface (GUI) automation. However, existing
GUI models mostly rely on learning from nearly error-free offline trajectories,
thus lacking reflection and error recovery capabilities. To bridge this gap, we
propose GUI-Reflection, a novel framework that explicitly integrates
self-reflection and error correction capabilities into end-to-end multimodal
GUI models throughout dedicated training stages: GUI-specific pre-training,
offline supervised fine-tuning (SFT), and online reflection tuning.
GUI-reflection enables self-reflection behavior emergence with fully automated
data generation and learning processes without requiring any human annotation.
Specifically, 1) we first propose scalable data pipelines to automatically
construct reflection and error correction data from existing successful
trajectories. While existing GUI models mainly focus on grounding and UI
understanding ability, we propose the GUI-Reflection Task Suite to learn and
evaluate reflection-oriented abilities explicitly. 2) Furthermore, we built a
diverse and efficient environment for online training and data collection of
GUI models on mobile devices. 3) We also present an iterative online reflection
tuning algorithm leveraging the proposed environment, enabling the model to
continuously enhance its reflection and error correction abilities. Our
framework equips GUI agents with self-reflection and correction capabilities,
paving the way for more robust, adaptable, and intelligent GUI automation, with
all data, models, environments, and tools to be released publicly.


## AI 摘要

该研究提出GUI-Reflection框架，通过三个阶段（预训练、监督微调、在线反思调优）赋予多模态GUI模型自我反思和纠错能力。其创新点包括：1) 自动从成功轨迹生成反思/纠错数据，并设计专门的任务套件；2) 搭建移动端在线训练环境；3) 开发迭代式在线调优算法。该方法无需人工标注，使GUI代理具备错误恢复能力，推动更鲁棒、自适应的GUI自动化发展。所有数据、模型和环境工具将开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T02:31:03Z
- **目录日期**: 2025-06-11

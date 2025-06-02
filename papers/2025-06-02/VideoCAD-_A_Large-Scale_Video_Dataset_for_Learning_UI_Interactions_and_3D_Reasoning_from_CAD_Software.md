# VideoCAD: A Large-Scale Video Dataset for Learning UI Interactions and 3D Reasoning from CAD Software

**URL**: http://arxiv.org/abs/2505.24838v1

## 原始摘要

Computer-Aided Design (CAD) is a time-consuming and complex process,
requiring precise, long-horizon user interactions with intricate 3D interfaces.
While recent advances in AI-driven user interface (UI) agents show promise,
most existing datasets and methods focus on short, low-complexity tasks in
mobile or web applications, failing to capture the demands of professional
engineering tools. In this work, we introduce VideoCAD, the first attempt at
engineering UI interaction learning for precision tasks. Specifically, VideoCAD
is a large-scale synthetic dataset consisting of over 41K annotated video
recordings of CAD operations, generated using an automated framework for
collecting high-fidelity UI action data from human-made CAD designs. Compared
to existing datasets, VideoCAD offers an order of magnitude higher complexity
in UI interaction learning for real-world engineering tasks, having up to a 20x
longer time horizon than other datasets. We show two important downstream
applications of VideoCAD: learning UI interactions from professional precision
3D CAD tools and a visual question-answering (VQA) benchmark designed to
evaluate multimodal large language models' (LLM) spatial reasoning and video
understanding abilities. To learn the UI interactions, we propose
VideoCADFormer - a state-of-the-art model in learning CAD interactions directly
from video, which outperforms multiple behavior cloning baselines. Both
VideoCADFormer and the VQA benchmark derived from VideoCAD reveal key
challenges in the current state of video-based UI understanding, including the
need for precise action grounding, multi-modal and spatial reasoning, and
long-horizon dependencies.


## AI 摘要

本文介绍了VideoCAD，首个针对专业工程工具（如CAD）的大规模合成数据集，包含41K带标注的CAD操作视频，其交互复杂度远超现有移动端/网页数据集（时间跨度达20倍）。研究提出VideoCADFormer模型，在视频级CAD交互学习中表现优于基线方法，并构建了多模态VQA基准测试。该工作揭示了视频UI理解的关键挑战：精确动作定位、多模态空间推理和长时依赖关系，为AI驱动的高精度工程界面交互提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T05:01:55Z
- **目录日期**: 2025-06-02

# FALCONEye: Finding Answers and Localizing Content in ONE-hour-long videos with multi-modal LLMs

**URL**: http://arxiv.org/abs/2503.19850v1

## 原始摘要

Information retrieval in hour-long videos presents a significant challenge,
even for state-of-the-art Vision-Language Models (VLMs), particularly when the
desired information is localized within a small subset of frames. Long video
data presents challenges for VLMs due to context window limitations and the
difficulty of pinpointing frames containing the answer. Our novel video agent,
FALCONEye, combines a VLM and a Large Language Model (LLM) to search relevant
information along the video, and locate the frames with the answer. FALCONEye
novelty relies on 1) the proposed meta-architecture, which is better suited to
tackle hour-long videos compared to short video approaches in the
state-of-the-art; 2) a new efficient exploration algorithm to locate the
information using short clips, captions and answer confidence; and 3) our
state-of-the-art VLMs calibration analysis for the answer confidence. Our agent
is built over a small-size VLM and a medium-size LLM being accessible to run on
standard computational resources. We also release FALCON-Bench, a benchmark to
evaluate long (average &gt; 1 hour) Video Answer Search challenges, highlighting
the need for open-ended question evaluation. Our experiments show FALCONEye's
superior performance than the state-of-the-art in FALCON-Bench, and similar or
better performance in related benchmarks.


## AI 摘要

FALCONEye是一种新型视频代理系统，结合视觉语言模型(VLM)和大语言模型(LLM)，专门解决时长超过1小时的视频信息检索难题。其创新点包括：1)针对长视频优化的元架构；2)基于短片段、字幕和答案置信度的高效探索算法；3)先进的VLM校准分析。该系统能在标准计算资源上运行，并在新发布的FALCON-Bench长视频基准测试中表现优于现有技术，同时在相关测试中保持相当或更好的性能。该研究还发布了专注于长视频开放问题评估的FALCON-Bench基准数据集。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T00:02:06Z
- **目录日期**: 2025-03-27

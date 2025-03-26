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

FALCONEye是一种新型视频代理，结合视觉语言模型(VLM)和大语言模型(LLM)，专门解决长视频(>1小时)信息检索难题。其创新点包括：1)针对长视频优化的元架构；2)基于短片段、字幕和置信度的高效探索算法；3)先进的VLM置信度校准方法。该系统仅需中小规模模型，可在标准硬件运行。研究团队同时发布了FALCON-Bench长视频问答基准测试。实验表明，FALCONEye在FALCON-Bench上超越现有技术，在其他基准测试中表现相当或更优，有效解决了长视频中关键帧定位的挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T18:02:03Z
- **目录日期**: 2025-03-26

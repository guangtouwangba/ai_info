# OVO-Bench: How Far is Your Video-LLMs from Real-World Online Video Understanding?

**URL**: http://arxiv.org/abs/2501.05510v2

## 原始摘要

Temporal Awareness, the ability to reason dynamically based on the timestamp
when a question is raised, is the key distinction between offline and online
video LLMs. Unlike offline models, which rely on complete videos for static,
post hoc analysis, online models process video streams incrementally and
dynamically adapt their responses based on the timestamp at which the question
is posed. Despite its significance, temporal awareness has not been adequately
evaluated in existing benchmarks. To fill this gap, we present OVO-Bench
(Online-VideO-Benchmark), a novel video benchmark that emphasizes the
importance of timestamps for advanced online video understanding capability
benchmarking. OVO-Bench evaluates the ability of video LLMs to reason and
respond to events occurring at specific timestamps under three distinct
scenarios: (1) Backward tracing: trace back to past events to answer the
question. (2) Real-time understanding: understand and respond to events as they
unfold at the current timestamp. (3) Forward active responding: delay the
response until sufficient future information becomes available to answer the
question accurately. OVO-Bench comprises 12 tasks, featuring 644 unique videos
and approximately human-curated 2,800 fine-grained meta-annotations with
precise timestamps. We combine automated generation pipelines with human
curation. With these high-quality samples, we further developed an evaluation
pipeline to systematically query video LLMs along the video timeline.
Evaluations of nine Video-LLMs reveal that, despite advancements on traditional
benchmarks, current models struggle with online video understanding, showing a
significant gap compared to human agents. We hope OVO-Bench will drive progress
in video LLMs and inspire future research in online video reasoning. Our
benchmark and code can be accessed at https://github.com/JoeLeelyf/OVO-Bench.


## AI 摘要

OVO-Bench是一个新型视频基准测试，旨在评估视频大语言模型（LLMs）的时间感知能力——即根据问题时间戳动态推理的能力。该基准测试聚焦三种场景：回溯过去事件（Backward tracing）、实时理解当前事件（Real-time understanding）和延迟响应以等待未来信息（Forward active responding）。包含12项任务、644个视频和约2,800条带精确时间戳的人工标注数据。评测显示，尽管现有模型在传统基准上表现良好，但在在线视频理解方面仍显著落后于人类。该研究希望推动视频LLMs在动态时序推理方向的发展。代码与数据已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-29T13:06:16Z
- **目录日期**: 2025-03-29

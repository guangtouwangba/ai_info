# T2A-Feedback: Improving Basic Capabilities of Text-to-Audio Generation via Fine-grained AI Feedback

**URL**: http://arxiv.org/abs/2505.10561v1

## 原始摘要

Text-to-audio (T2A) generation has achieved remarkable progress in generating
a variety of audio outputs from language prompts. However, current
state-of-the-art T2A models still struggle to satisfy human preferences for
prompt-following and acoustic quality when generating complex multi-event
audio. To improve the performance of the model in these high-level
applications, we propose to enhance the basic capabilities of the model with AI
feedback learning. First, we introduce fine-grained AI audio scoring pipelines
to: 1) verify whether each event in the text prompt is present in the audio
(Event Occurrence Score), 2) detect deviations in event sequences from the
language description (Event Sequence Score), and 3) assess the overall acoustic
and harmonic quality of the generated audio (Acoustic&amp;Harmonic Quality). We
evaluate these three automatic scoring pipelines and find that they correlate
significantly better with human preferences than other evaluation metrics. This
highlights their value as both feedback signals and evaluation metrics.
Utilizing our robust scoring pipelines, we construct a large audio preference
dataset, T2A-FeedBack, which contains 41k prompts and 249k audios, each
accompanied by detailed scores. Moreover, we introduce T2A-EpicBench, a
benchmark that focuses on long captions, multi-events, and story-telling
scenarios, aiming to evaluate the advanced capabilities of T2A models. Finally,
we demonstrate how T2A-FeedBack can enhance current state-of-the-art audio
model. With simple preference tuning, the audio generation model exhibits
significant improvements in both simple (AudioCaps test set) and complex
(T2A-EpicBench) scenarios.


## AI 摘要

该研究提出通过AI反馈学习提升文本到音频(T2A)生成模型在复杂多事件音频中的表现。研究者开发了细粒度音频评分系统，评估事件完整性、序列准确性和音质，并构建了包含41k提示和249k音频的T2A-FeedBack数据集。同时推出T2A-EpicBench基准测试，专注长文本和多事件场景。实验表明，这些评分指标与人类偏好显著相关，且通过简单偏好调优即可显著提升现有模型在简单和复杂场景下的表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-18T06:01:01Z
- **目录日期**: 2025-05-18

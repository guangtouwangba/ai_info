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

本文提出了一种基于AI反馈学习的文本到音频（T2A）生成改进方法。研究者开发了三个细粒度评分管道：事件发生评分、事件序列评分和声学质量评分，这些评分与人类偏好显著相关。基于此，构建了包含41k提示和249k音频的大规模偏好数据集T2A-FeedBack，并提出了针对长文本、多事件场景的评测基准T2A-EpicBench。实验表明，通过简单的偏好调优，现有音频生成模型在简单（AudioCaps）和复杂（T2A-EpicBench）场景下均获得显著提升。该方法有效提高了模型在复杂多事件音频生成中的表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-18T14:01:00Z
- **目录日期**: 2025-05-18

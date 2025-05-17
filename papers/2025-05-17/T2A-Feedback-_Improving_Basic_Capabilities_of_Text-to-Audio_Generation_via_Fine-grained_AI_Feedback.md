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

当前文本转音频(T2A)模型在处理复杂多事件音频时仍存在提示跟随和音质问题。研究提出通过AI反馈学习提升模型能力：1)开发细粒度评分系统，评估事件完整性、顺序准确性和音质；2)构建含4.1万提示和24.9万音频的T2A-FeedBack数据集；3)推出针对长文本、多事件场景的T2A-EpicBench基准测试。实验表明，该评分系统与人类偏好高度相关。通过简单偏好调优，模型在简单(AudioCaps)和复杂(T2A-EpicBench)场景下均显著提升生成质量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T16:01:02Z
- **目录日期**: 2025-05-17

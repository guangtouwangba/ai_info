# Towards Online Multi-Modal Social Interaction Understanding

**URL**: http://arxiv.org/abs/2503.19851v1

## 原始摘要

Multimodal social interaction understanding (MMSI) is critical in human-robot
interaction systems. In real-world scenarios, AI agents are required to provide
real-time feedback. However, existing models often depend on both past and
future contexts, which hinders them from applying to real-world problems. To
bridge this gap, we propose an online MMSI setting, where the model must
resolve MMSI tasks using only historical information, such as recorded
dialogues and video streams. To address the challenges of missing the useful
future context, we develop a novel framework, named Online-MMSI-VLM, that
leverages two complementary strategies: multi-party conversation forecasting
and social-aware visual prompting with multi-modal large language models.
First, to enrich linguistic context, the multi-party conversation forecasting
simulates potential future utterances in a coarse-to-fine manner, anticipating
upcoming speaker turns and then generating fine-grained conversational details.
Second, to effectively incorporate visual social cues like gaze and gesture,
social-aware visual prompting highlights the social dynamics in video with
bounding boxes and body keypoints for each person and frame. Extensive
experiments on three tasks and two datasets demonstrate that our method
achieves state-of-the-art performance and significantly outperforms baseline
models, indicating its effectiveness on Online-MMSI. The code and pre-trained
models will be publicly released at: https://github.com/Sampson-Lee/OnlineMMSI.


## AI 摘要

本文提出了一种在线多模态社交交互理解框架Online-MMSI-VLM，解决了现有模型依赖未来上下文而无法实时应用的局限性。该框架采用两种互补策略：1）多轮对话预测，通过粗粒度（预测说话人）到细粒度（生成对话内容）的方式模拟未来对话；2）社会感知视觉提示，利用边界框和人体关键点突出视频中的社交动态（如眼神和手势）。在三个任务和两个数据集上的实验表明，该方法显著优于基线模型，达到最先进性能。代码和预训练模型将开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T09:04:03Z
- **目录日期**: 2025-03-26

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

该研究提出了一种在线多模态社交交互理解（Online-MMSI）框架，用于解决实时人机交互中未来上下文缺失的问题。该框架采用两种互补策略：1）通过从粗到细的多方对话预测模拟未来可能的对话内容；2）利用社交感知视觉提示（如边界框和人体关键点）突出视频中的社交动态。实验表明，该方法在三个任务和两个数据集上均优于基线模型，达到最先进水平。代码和预训练模型将开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T23:09:06Z
- **目录日期**: 2025-03-26

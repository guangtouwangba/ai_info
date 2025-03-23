# GAEA: A Geolocation Aware Conversational Model

**URL**: http://arxiv.org/abs/2503.16423v1

## 原始摘要

Image geolocalization, in which, traditionally, an AI model predicts the
precise GPS coordinates of an image is a challenging task with many downstream
applications. However, the user cannot utilize the model to further their
knowledge other than the GPS coordinate; the model lacks an understanding of
the location and the conversational ability to communicate with the user. In
recent days, with tremendous progress of large multimodal models (LMMs)
proprietary and open-source researchers have attempted to geolocalize images
via LMMs. However, the issues remain unaddressed; beyond general tasks, for
more specialized downstream tasks, one of which is geolocalization, LMMs
struggle. In this work, we propose to solve this problem by introducing a
conversational model GAEA that can provide information regarding the location
of an image, as required by a user. No large-scale dataset enabling the
training of such a model exists. Thus we propose a comprehensive dataset GAEA
with 800K images and around 1.6M question answer pairs constructed by
leveraging OpenStreetMap (OSM) attributes and geographical context clues. For
quantitative evaluation, we propose a diverse benchmark comprising 4K
image-text pairs to evaluate conversational capabilities equipped with diverse
question types. We consider 11 state-of-the-art open-source and proprietary
LMMs and demonstrate that GAEA significantly outperforms the best open-source
model, LLaVA-OneVision by 25.69% and the best proprietary model, GPT-4o by
8.28%. Our dataset, model and codes are available


## AI 摘要

图像地理定位是一项具有挑战性的任务，传统AI模型仅能预测图像的GPS坐标，缺乏对位置的理解和与用户的对话能力。为解决这一问题，研究者提出了GAEA模型，能够根据用户需求提供图像位置信息。由于缺乏大规模数据集，研究者构建了包含80万张图像和160万问答对的GAEA数据集，利用OpenStreetMap属性和地理上下文线索。通过包含4000个图像-文本对的多样化基准测试，GAEA在对话能力上显著优于现有开源和专有模型，分别比LLaVA-OneVision和GPT-4o高出25.69%和8.28%。数据集、模型和代码均已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-23T15:00:59Z
- **目录日期**: 2025-03-23

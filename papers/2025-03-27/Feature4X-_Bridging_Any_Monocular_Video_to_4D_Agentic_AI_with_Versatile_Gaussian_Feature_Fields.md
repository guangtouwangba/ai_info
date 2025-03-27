# Feature4X: Bridging Any Monocular Video to 4D Agentic AI with Versatile Gaussian Feature Fields

**URL**: http://arxiv.org/abs/2503.20776v1

## 原始摘要

Recent advancements in 2D and multimodal models have achieved remarkable
success by leveraging large-scale training on extensive datasets. However,
extending these achievements to enable free-form interactions and high-level
semantic operations with complex 3D/4D scenes remains challenging. This
difficulty stems from the limited availability of large-scale, annotated 3D/4D
or multi-view datasets, which are crucial for generalizable vision and language
tasks such as open-vocabulary and prompt-based segmentation, language-guided
editing, and visual question answering (VQA). In this paper, we introduce
Feature4X, a universal framework designed to extend any functionality from 2D
vision foundation model into the 4D realm, using only monocular video input,
which is widely available from user-generated content. The "X" in Feature4X
represents its versatility, enabling any task through adaptable,
model-conditioned 4D feature field distillation. At the core of our framework
is a dynamic optimization strategy that unifies multiple model capabilities
into a single representation. Additionally, to the best of our knowledge,
Feature4X is the first method to distill and lift the features of video
foundation models (e.g. SAM2, InternVideo2) into an explicit 4D feature field
using Gaussian Splatting. Our experiments showcase novel view segment anything,
geometric and appearance scene editing, and free-form VQA across all time
steps, empowered by LLMs in feedback loops. These advancements broaden the
scope of agentic AI applications by providing a foundation for scalable,
contextually and spatiotemporally aware systems capable of immersive dynamic 4D
scene interaction.


## AI 摘要

Feature4X是一个通用框架，可将2D视觉基础模型的功能扩展到4D领域，仅需单目视频输入。它通过动态优化策略，将多种模型能力统一到单一表示中，并首次利用高斯泼溅（Gaussian Splatting）将视频基础模型（如SAM2、InternVideo2）的特征提取为显式4D特征场。该框架支持新颖视角分割、场景几何与外观编辑，以及跨时间步的自由形式视觉问答（VQA），结合大语言模型（LLM）实现反馈循环。Feature4X为可扩展、具有时空感知的沉浸式动态4D场景交互系统奠定了基础，推动了智能体AI应用的发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T07:04:06Z
- **目录日期**: 2025-03-27

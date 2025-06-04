# GUI-Actor: Coordinate-Free Visual Grounding for GUI Agents

**URL**: http://arxiv.org/abs/2506.03143v1

## 原始摘要

One of the principal challenges in building VLM-powered GUI agents is visual
grounding, i.e., localizing the appropriate screen region for action execution
based on both the visual content and the textual plans. Most existing work
formulates this as a text-based coordinate generation task. However, these
approaches suffer from several limitations: weak spatial-semantic alignment,
inability to handle ambiguous supervision targets, and a mismatch between the
dense nature of screen coordinates and the coarse, patch-level granularity of
visual features extracted by models like Vision Transformers. In this paper, we
propose GUI-Actor, a VLM-based method for coordinate-free GUI grounding. At its
core, GUI-Actor introduces an attention-based action head that learns to align
a dedicated <actor> token with all relevant visual patch tokens, enabling the
model to propose one or more action regions in a single forward pass. In line
with this, we further design a grounding verifier to evaluate and select the
most plausible action region from the candidates proposed for action execution.
Extensive experiments show that GUI-Actor outperforms prior state-of-the-art
methods on multiple GUI action grounding benchmarks, with improved
generalization to unseen screen resolutions and layouts. Notably, GUI-Actor-7B
even surpasses UI-TARS-72B (38.1) on ScreenSpot-Pro, achieving scores of 40.7
with Qwen2-VL and 44.6 with Qwen2.5-VL as backbones. Furthermore, by
incorporating the verifier, we find that fine-tuning only the newly introduced
action head (~100M parameters for 7B model) while keeping the VLM backbone
frozen is sufficient to achieve performance comparable to previous
state-of-the-art models, highlighting that GUI-Actor can endow the underlying
VLM with effective grounding capabilities without compromising its
general-purpose strengths.
</actor>

## AI 摘要

本文提出GUI-Actor，一种基于视觉语言模型（VLM）的无坐标GUI视觉定位方法。该方法通过引入注意力机制的动作头，将专用<actor>标记与相关视觉块对齐，实现单次前向传递即可定位多个动作区域，并设计了验证器筛选最优执行区域。实验表明，GUI-Actor在多个基准测试中超越现有方法，尤其在跨分辨率/布局泛化性上表现突出。例如，7B参数的GUI-Actor在ScreenSpot-Pro上以40.7/44.6分超越72B参数的UI-TARS（38.1分）。仅微调动作头（约1亿参数）即可达到SOTA性能，证明该方法能在保留VLM通用能力的同时赋予精准定位能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T12:01:26Z
- **目录日期**: 2025-06-04

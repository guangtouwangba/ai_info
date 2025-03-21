# Do Visual Imaginations Improve Vision-and-Language Navigation Agents?

**URL**: http://arxiv.org/abs/2503.16394v1

## 原始摘要

Vision-and-Language Navigation (VLN) agents are tasked with navigating an
unseen environment using natural language instructions. In this work, we study
if visual representations of sub-goals implied by the instructions can serve as
navigational cues and lead to increased navigation performance. To synthesize
these visual representations or imaginations, we leverage a text-to-image
diffusion model on landmark references contained in segmented instructions.
These imaginations are provided to VLN agents as an added modality to act as
landmark cues and an auxiliary loss is added to explicitly encourage relating
these with their corresponding referring expressions. Our findings reveal an
increase in success rate (SR) of around 1 point and up to 0.5 points in success
scaled by inverse path length (SPL) across agents. These results suggest that
the proposed approach reinforces visual understanding compared to relying on
language instructions alone. Code and data for our work can be found at
https://www.akhilperincherry.com/VLN-Imagine-website/.


## AI 摘要

本研究探讨了在视觉与语言导航（VLN）任务中，通过文本到图像扩散模型生成子目标的视觉表示（即“想象”）作为导航线索的效果。这些视觉表示基于分段指令中的地标参考生成，并作为额外模态提供给VLN代理，同时引入辅助损失以强化其与对应语言表达的联系。实验结果表明，该方法显著提升了导航成功率（SR）约1个百分点，路径长度加权成功率（SPL）提升最高达0.5个百分点，表明视觉想象增强了代理的视觉理解能力，优于仅依赖语言指令的导航策略。相关代码和数据已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-21T13:07:17Z
- **目录日期**: 2025-03-21

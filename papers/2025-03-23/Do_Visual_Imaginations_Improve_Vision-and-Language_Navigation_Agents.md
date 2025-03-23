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

本研究探讨了在视觉与语言导航（VLN）任务中，利用文本到图像扩散模型生成指令中隐含子目标的视觉表示，作为导航线索以提高导航性能。通过将生成的视觉想象作为额外模态提供给VLN代理，并引入辅助损失函数以强化视觉与语言表达的关联，实验结果显示，代理的成功率（SR）提高了约1个百分点，路径长度加权的成功率（SPL）提高了最多0.5个百分点。这表明，相较于仅依赖语言指令，该方法增强了视觉理解能力。相关代码和数据可在指定网站获取。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-23T13:05:20Z
- **目录日期**: 2025-03-23

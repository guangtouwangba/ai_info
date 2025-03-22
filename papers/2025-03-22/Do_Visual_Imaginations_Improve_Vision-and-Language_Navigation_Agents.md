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

本研究探讨了在视觉与语言导航（VLN）任务中，利用自然语言指令中隐含的子目标视觉表示作为导航线索是否能提升导航性能。通过文本到图像的扩散模型生成这些视觉表示，并将其作为额外模态提供给VLN代理，以作为地标线索。同时，引入辅助损失函数以强化这些视觉表示与对应语言表达的联系。实验结果表明，该方法在成功率（SR）和路径长度调整成功率（SPL）上分别提升了约1点和0.5点，表明相比仅依赖语言指令，该方法增强了视觉理解能力。代码和数据可在指定网站获取。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-22T06:01:43Z
- **目录日期**: 2025-03-22

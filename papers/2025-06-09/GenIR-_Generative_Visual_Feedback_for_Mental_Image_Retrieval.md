# GenIR: Generative Visual Feedback for Mental Image Retrieval

**URL**: http://arxiv.org/abs/2506.06220v1

## 原始摘要

Vision-language models (VLMs) have shown strong performance on text-to-image
retrieval benchmarks. However, bridging this success to real-world applications
remains a challenge. In practice, human search behavior is rarely a one-shot
action. Instead, it is often a multi-round process guided by clues in mind,
that is, a mental image ranging from vague recollections to vivid mental
representations of the target image. Motivated by this gap, we study the task
of Mental Image Retrieval (MIR), which targets the realistic yet underexplored
setting where users refine their search for a mentally envisioned image through
multi-round interactions with an image search engine. Central to successful
interactive retrieval is the capability of machines to provide users with
clear, actionable feedback; however, existing methods rely on indirect or
abstract verbal feedback, which can be ambiguous, misleading, or ineffective
for users to refine the query. To overcome this, we propose GenIR, a generative
multi-round retrieval paradigm leveraging diffusion-based image generation to
explicitly reify the AI system's understanding at each round. These synthetic
visual representations provide clear, interpretable feedback, enabling users to
refine their queries intuitively and effectively. We further introduce a fully
automated pipeline to generate a high-quality multi-round MIR dataset.
Experimental results demonstrate that GenIR significantly outperforms existing
interactive methods in the MIR scenario. This work establishes a new task with
a dataset and an effective generative retrieval method, providing a foundation
for future research in this direction.


## AI 摘要

本文提出"心理图像检索"(MIR)任务，研究用户通过多轮交互在脑海中寻找目标图像的现实场景。针对现有方法依赖模糊抽象反馈的问题，作者开发了GenIR生成式检索框架，利用扩散模型在每轮交互生成可视化反馈图像，帮助用户直观调整查询。该研究还构建了高质量多轮MIR数据集。实验表明，GenIR显著优于现有交互方法，为这一新兴方向奠定了基础。这项工作通过引入新任务、数据集和生成式检索方法，推动了视觉语言模型在真实场景中的应用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T10:01:38Z
- **目录日期**: 2025-06-09

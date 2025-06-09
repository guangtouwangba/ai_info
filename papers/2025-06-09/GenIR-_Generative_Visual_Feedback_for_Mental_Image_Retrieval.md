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

本文提出了一种基于生成式交互的图像检索方法GenIR，用于解决用户通过多轮交互搜索脑海中模糊图像（Mental Image Retrieval, MIR）的挑战。传统方法依赖抽象文字反馈，效果有限；而GenIR利用扩散模型在每轮交互中生成可视化反馈图像，使用户能直观调整查询。研究还构建了自动化生成的高质量多轮MIR数据集。实验表明，GenIR显著优于现有交互式检索方法。该工作不仅定义了新任务，还提供了数据集和有效方法，为未来相关研究奠定了基础。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T13:11:21Z
- **目录日期**: 2025-06-09

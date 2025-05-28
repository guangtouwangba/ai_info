# Paper2Poster: Towards Multimodal Poster Automation from Scientific Papers

**URL**: http://arxiv.org/abs/2505.21497v1

## 原始摘要

Academic poster generation is a crucial yet challenging task in scientific
communication, requiring the compression of long-context interleaved documents
into a single, visually coherent page. To address this challenge, we introduce
the first benchmark and metric suite for poster generation, which pairs recent
conference papers with author-designed posters and evaluates outputs on
(i)Visual Quality-semantic alignment with human posters, (ii)Textual
Coherence-language fluency, (iii)Holistic Assessment-six fine-grained aesthetic
and informational criteria scored by a VLM-as-judge, and notably
(iv)PaperQuiz-the poster's ability to convey core paper content as measured by
VLMs answering generated quizzes. Building on this benchmark, we propose
PosterAgent, a top-down, visual-in-the-loop multi-agent pipeline: the (a)Parser
distills the paper into a structured asset library; the (b)Planner aligns
text-visual pairs into a binary-tree layout that preserves reading order and
spatial balance; and the (c)Painter-Commenter loop refines each panel by
executing rendering code and using VLM feedback to eliminate overflow and
ensure alignment. In our comprehensive evaluation, we find that GPT-4o
outputs-though visually appealing at first glance-often exhibit noisy text and
poor PaperQuiz scores, and we find that reader engagement is the primary
aesthetic bottleneck, as human-designed posters rely largely on visual
semantics to convey meaning. Our fully open-source variants (e.g. based on the
Qwen-2.5 series) outperform existing 4o-driven multi-agent systems across
nearly all metrics, while using 87% fewer tokens. It transforms a 22-page paper
into a finalized yet editable .pptx poster - all for just $0.005. These
findings chart clear directions for the next generation of fully automated
poster-generation models. The code and datasets are available at
https://github.com/Paper2Poster/Paper2Poster.


## AI 摘要

本研究提出了首个学术海报生成基准和评估体系，包含视觉质量、文本连贯性、整体评估和论文内容传达能力四个维度。基于此，团队开发了PosterAgent系统，采用自上而下的多智能体流程：解析器提取论文结构、规划器设计布局、绘制-评论循环优化细节。实验表明，GPT-4o生成的海报虽视觉吸引但文本质量差，而开源模型Qwen-2.5系列在多数指标上表现更优，且成本降低87%。该系统能以0.005美元成本将22页论文转换为可编辑.pptx海报，为自动化海报生成指明方向。代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T13:11:20Z
- **目录日期**: 2025-05-28

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

本研究提出了首个学术海报生成基准和评估指标，开发了PosterAgent多智能体流程。该系统通过解析论文、规划布局和循环优化三个步骤，将22页论文转化为可编辑的.pptx海报，成本仅0.005美元。相比GPT-4o生成的海报（虽美观但文本质量差），开源版本(Qwen-2.5系列)在各项指标上表现更优，且减少87%的token使用。评估发现人类设计海报更依赖视觉语义传达信息，而读者参与度是美学瓶颈。该研究为全自动海报生成模型指明了方向，相关代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T23:01:27Z
- **目录日期**: 2025-05-28

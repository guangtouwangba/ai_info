# VBench-2.0: Advancing Video Generation Benchmark Suite for Intrinsic Faithfulness

**URL**: http://arxiv.org/abs/2503.21755v1

## 原始摘要

Video generation has advanced significantly, evolving from producing
unrealistic outputs to generating videos that appear visually convincing and
temporally coherent. To evaluate these video generative models, benchmarks such
as VBench have been developed to assess their faithfulness, measuring factors
like per-frame aesthetics, temporal consistency, and basic prompt adherence.
However, these aspects mainly represent superficial faithfulness, which focus
on whether the video appears visually convincing rather than whether it adheres
to real-world principles. While recent models perform increasingly well on
these metrics, they still struggle to generate videos that are not just
visually plausible but fundamentally realistic. To achieve real "world models"
through video generation, the next frontier lies in intrinsic faithfulness to
ensure that generated videos adhere to physical laws, commonsense reasoning,
anatomical correctness, and compositional integrity. Achieving this level of
realism is essential for applications such as AI-assisted filmmaking and
simulated world modeling. To bridge this gap, we introduce VBench-2.0, a
next-generation benchmark designed to automatically evaluate video generative
models for their intrinsic faithfulness. VBench-2.0 assesses five key
dimensions: Human Fidelity, Controllability, Creativity, Physics, and
Commonsense, each further broken down into fine-grained capabilities. Tailored
for individual dimensions, our evaluation framework integrates generalists such
as state-of-the-art VLMs and LLMs, and specialists, including anomaly detection
methods proposed for video generation. We conduct extensive annotations to
ensure alignment with human judgment. By pushing beyond superficial
faithfulness toward intrinsic faithfulness, VBench-2.0 aims to set a new
standard for the next generation of video generative models in pursuit of
intrinsic faithfulness.


## AI 摘要

视频生成技术已从早期不真实的输出发展到能生成视觉逼真且时间连贯的视频。现有评估基准（如VBench）主要关注表面真实性（如帧美学、时间一致性等），但忽略了内在真实性（如物理法则、常识推理等）。为此，研究者提出VBench-2.0，一个新一代评估基准，自动检测视频生成模型的内在真实性，涵盖人类逼真度、可控性、创造力、物理法则和常识五大维度，并细分为更精细的能力。该框架结合通用模型（如VLM、LLM）和专用方法（如异常检测），通过大量标注确保与人类判断一致，旨在为下一代视频生成模型设定新标准。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-28T04:01:11Z
- **目录日期**: 2025-03-28

# Kimi-VL Technical Report

**URL**: http://arxiv.org/abs/2504.07491v2

## 原始摘要

We present Kimi-VL, an efficient open-source Mixture-of-Experts (MoE)
vision-language model (VLM) that offers advanced multimodal reasoning,
long-context understanding, and strong agent capabilities - all while
activating only 2.8B parameters in its language decoder (Kimi-VL-A3B). Kimi-VL
demonstrates strong performance across challenging domains: as a
general-purpose VLM, Kimi-VL excels in multi-turn agent tasks (e.g., OSWorld),
matching flagship models. Furthermore, it exhibits remarkable capabilities
across diverse challenging vision language tasks, including college-level image
and video comprehension, OCR, mathematical reasoning, and multi-image
understanding. In comparative evaluations, it effectively competes with
cutting-edge efficient VLMs such as GPT-4o-mini, Qwen2.5-VL-7B, and
Gemma-3-12B-IT, while surpassing GPT-4o in several key domains. Kimi-VL also
advances in processing long contexts and perceiving clearly. With a 128K
extended context window, Kimi-VL can process diverse long inputs, achieving
impressive scores of 64.5 on LongVideoBench and 35.1 on MMLongBench-Doc. Its
native-resolution vision encoder, MoonViT, further allows it to see and
understand ultra-high-resolution visual inputs, achieving 83.2 on InfoVQA and
34.5 on ScreenSpot-Pro, while maintaining lower computational cost for common
tasks. Building upon Kimi-VL, we introduce an advanced long-thinking variant:
Kimi-VL-Thinking. Developed through long chain-of-thought (CoT) supervised
fine-tuning (SFT) and reinforcement learning (RL), this model exhibits strong
long-horizon reasoning capabilities. It achieves scores of 61.7 on MMMU, 36.8
on MathVision, and 71.3 on MathVista while maintaining the compact 2.8B
activated LLM parameters, setting a new standard for efficient multimodal
thinking models. Code and models are publicly accessible at
https://github.com/MoonshotAI/Kimi-VL.


## AI 摘要

Kimi-VL是一款高效的开源视觉语言模型（VLM），采用混合专家（MoE）架构，仅激活28亿参数的语言解码器。它在多模态推理、长上下文理解和智能体任务中表现优异，在OSWorld等多轮任务中媲美顶级模型，并在图像/视频理解、OCR、数学推理等任务中超越GPT-4o等先进模型。支持128K长上下文窗口，在LongVideoBench和MMLongBench-Doc上分别获得64.5和35.1分。其原生分辨率视觉编码器MoonViT能处理超高清输入，在InfoVQA和ScreenSpot-Pro上达83.2和34.5分。进阶版Kimi-VL-Thinking通过长链思维微调，在MMMU等数学任务中刷新高效多模态模型的标杆。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T13:10:33Z
- **目录日期**: 2025-04-16

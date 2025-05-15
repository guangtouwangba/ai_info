# Customizing a Large Language Model for VHDL Design of High-Performance Microprocessors

**URL**: http://arxiv.org/abs/2505.09610v1

## 原始摘要

The use of Large Language Models (LLMs) in hardware design has taken off in
recent years, principally through its incorporation in tools that increase chip
designer productivity. There has been considerable discussion about the use of
LLMs in RTL specifications of chip designs, for which the two most popular
languages are Verilog and VHDL. LLMs and their use in Verilog design has
received significant attention due to the higher popularity of the language,
but little attention so far has been given to VHDL despite its continued
popularity in the industry. There has also been little discussion about the
unique needs of organizations that engage in high-performance processor design,
and techniques to deploy AI solutions in these settings. In this paper, we
describe our journey in developing a Large Language Model (LLM) specifically
for the purpose of explaining VHDL code, a task that has particular importance
in an organization with decades of experience and assets in high-performance
processor design. We show how we developed test sets specific to our needs and
used them for evaluating models as we performed extended pretraining (EPT) of a
base LLM. Expert evaluation of the code explanations produced by the EPT model
increased to 69% compared to a base model rating of 43%. We further show how we
developed an LLM-as-a-judge to gauge models similar to expert evaluators. This
led us to deriving and evaluating a host of new models, including an
instruction-tuned version of the EPT model with an expected expert evaluator
rating of 71%. Our experiments also indicate that with the potential use of
newer base models, this rating can be pushed to 85% and beyond. We conclude
with a discussion on further improving the quality of hardware design LLMs
using exciting new developments in the Generative AI world.


## AI 摘要

近年来，大型语言模型(LLMs)在硬件设计中得到应用，主要提升芯片设计效率。虽然Verilog设计中的LLM应用受到关注，但VHDL领域的研究较少，尤其在高性能处理器设计方面。本文介绍了开发专用于VHDL代码解释的LLM的过程，包括创建定制测试集和进行扩展预训练(EPT)。专家评估显示EPT模型的代码解释准确率从基础模型的43%提升至69%。研究还开发了LLM-as-a-judge评估方法，指令调优后的EPT模型预期准确率达71%。实验表明，使用更新的基础模型可将准确率提升至85%以上。最后探讨了利用生成式AI新进展进一步提升硬件设计LLM质量的可能性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T22:00:58Z
- **目录日期**: 2025-05-15

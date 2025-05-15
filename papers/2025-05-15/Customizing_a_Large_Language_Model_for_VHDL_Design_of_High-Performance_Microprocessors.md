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

近年来，大型语言模型（LLM）在硬件设计中的应用逐渐兴起，主要用于提升芯片设计效率。虽然Verilog设计中的LLM应用受到较多关注，但VHDL语言的应用研究较少，尽管其在工业界仍广泛使用。本文探讨了针对高性能处理器设计机构的需求，开发专门用于解释VHDL代码的LLM的过程。通过定制测试集和扩展预训练（EPT），模型生成的代码解释专家评分从基础模型的43%提升至69%。研究还开发了LLM-as-a-judge评估方法，指令调优后的EPT模型预期评分达71%，未来使用新基础模型有望提升至85%以上。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T07:01:00Z
- **目录日期**: 2025-05-15

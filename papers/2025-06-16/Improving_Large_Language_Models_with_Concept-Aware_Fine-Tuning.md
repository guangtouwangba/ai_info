# Improving Large Language Models with Concept-Aware Fine-Tuning

**URL**: http://arxiv.org/abs/2506.07833v2

## 原始摘要

Large language models (LLMs) have become the cornerstone of modern AI.
However, the existing paradigm of next-token prediction fundamentally limits
their ability to form coherent, high-level concepts, making it a critical
barrier to human-like understanding and reasoning. Take the phrase "ribonucleic
acid" as an example: an LLM will first decompose it into tokens, i.e.,
artificial text fragments ("rib", "on", ...), then learn each token
sequentially, rather than grasping the phrase as a unified, coherent semantic
entity. This fragmented representation hinders deeper conceptual understanding
and, ultimately, the development of truly intelligent systems. In response, we
introduce Concept-Aware Fine-Tuning (CAFT), a novel multi-token training method
that redefines how LLMs are fine-tuned. By enabling the learning of sequences
that span multiple tokens, this method fosters stronger concept-aware learning.
Our experiments demonstrate significant improvements compared to conventional
next-token finetuning methods across diverse tasks, including traditional
applications like text summarization and domain-specific ones like de novo
protein design. Multi-token prediction was previously only possible in the
prohibitively expensive pretraining phase; CAFT, to our knowledge, is the first
to bring the multi-token setting to the post-training phase, thus effectively
democratizing its benefits for the broader community of practitioners and
researchers. Finally, the unexpected effectiveness of our proposed method
suggests wider implications for the machine learning research community. All
code and data are available at https://github.com/michaelchen-lab/caft-llm


## AI 摘要

这篇论文提出了概念感知微调(CAFT)方法，解决了大语言模型(LLM)在传统逐词预测范式下难以形成连贯高层概念的局限性。CAFT通过多词序列学习，使模型能整体理解语义实体(如"核糖核酸")而非碎片化token。实验表明，该方法在文本摘要、蛋白质设计等任务上显著优于传统微调方式。创新性地将原本仅适用于预训练阶段的多词预测扩展到了微调阶段，降低了应用门槛。这一发现对机器学习研究具有广泛启示意义。相关代码和数据已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T02:34:30Z
- **目录日期**: 2025-06-16

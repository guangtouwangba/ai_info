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

这篇论文提出了一种名为"概念感知微调"(CAFT)的新方法，解决了大型语言模型(LLMs)在传统"下一个token预测"范式下的局限性。现有方法将输入分解为独立token(如将"ribonucleic acid"拆分为"rib"、"on"等)，阻碍了模型对整体概念的连贯理解。CAFT通过多token序列训练，使模型能学习跨token的概念表示。实验表明，该方法在文本摘要、蛋白质设计等任务中显著优于传统微调方法。CAFT首次将多token预测从昂贵的预训练阶段引入微调阶段，为研究社区提供了更易获得的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T03:26:39Z
- **目录日期**: 2025-06-16

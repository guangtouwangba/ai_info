# OneProt: Towards Multi-Modal Protein Foundation Models

**URL**: http://arxiv.org/abs/2411.04863v2

## 原始摘要

Recent advances in Artificial Intelligence have enabled multi-modal systems
to model and translate diverse information spaces. Extending beyond text and
vision, we introduce OneProt, a multi-modal AI for proteins that integrates
structural, sequence, text, and binding site data. Using the ImageBind
framework, OneProt aligns the latent spaces of protein modality encoders in a
lightweight fine-tuning scheme that focuses on pairwise alignment with sequence
data rather than requiring full matches. This novel approach comprises a mix of
Graph Neural Networks and transformer architectures. It demonstrates strong
performance in retrieval tasks and showcases the efficacy of multi-modal
systems in Protein Machine Learning through a broad spectrum of downstream
baselines, including enzyme function prediction and binding site analysis.
Furthermore, OneProt enables the transfer of representational information from
specialized encoders to the sequence encoder, enhancing capabilities for
distinguishing evolutionarily related and unrelated sequences and exhibiting
representational properties where evolutionarily related proteins align in
similar directions within the latent space. In addition, we extensively
investigate modality ablations to identify the encoders that contribute most to
predictive performance, highlighting the significance of the binding site
encoder, which has not been used in similar models previously. This work
expands the horizons of multi-modal protein models, paving the way for
transformative applications in drug discovery, biocatalytic reaction planning,
and protein engineering.


## AI 摘要

研究者开发了名为OneProt的多模态AI系统，整合蛋白质的结构、序列、文本和结合位点数据。基于ImageBind框架，该系统采用图神经网络和Transformer架构，通过轻量级微调实现蛋白质模态编码器的潜在空间对齐，重点优化与序列数据的成对匹配而非完全匹配。实验表明，OneProt在检索任务及酶功能预测、结合位点分析等下游任务中表现优异，并能将专业编码器的表征信息迁移至序列编码器，增强进化相关/无关序列的区分能力。结合位点编码器的贡献尤为突出，这在同类模型中尚属首次应用。该研究为药物研发和蛋白质工程开辟了新途径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T21:02:17Z
- **目录日期**: 2025-05-26

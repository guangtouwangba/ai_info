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

研究人员开发了名为OneProt的多模态AI系统，可整合蛋白质的结构、序列、文本和结合位点数据。该系统基于ImageBind框架，采用图神经网络和Transformer架构，通过轻量级微调方案对齐不同蛋白质模态的潜在空间。OneProt在检索任务中表现优异，支持酶功能预测和结合位点分析等下游应用，并能将进化相关蛋白质在潜在空间中定向对齐。研究特别指出结合位点编码器对预测性能的重要贡献，这在以往类似模型中未被使用。该技术为药物发现、生物催化反应规划和蛋白质工程开辟了新途径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T23:02:22Z
- **目录日期**: 2025-05-26

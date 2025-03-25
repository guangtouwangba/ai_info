# FlowDock: Geometric Flow Matching for Generative Protein-Ligand Docking and Affinity Prediction

**URL**: http://arxiv.org/abs/2412.10966v3

## 原始摘要

Powerful generative AI models of protein-ligand structure have recently been
proposed, but few of these methods support both flexible protein-ligand docking
and affinity estimation. Of those that do, none can directly model multiple
binding ligands concurrently or have been rigorously benchmarked on
pharmacologically relevant drug targets, hindering their widespread adoption in
drug discovery efforts. In this work, we propose FlowDock, the first deep
geometric generative model based on conditional flow matching that learns to
directly map unbound (apo) structures to their bound (holo) counterparts for an
arbitrary number of binding ligands. Furthermore, FlowDock provides predicted
structural confidence scores and binding affinity values with each of its
generated protein-ligand complex structures, enabling fast virtual screening of
new (multi-ligand) drug targets. For the well-known PoseBusters Benchmark
dataset, FlowDock outperforms single-sequence AlphaFold 3 with a 51% blind
docking success rate using unbound (apo) protein input structures and without
any information derived from multiple sequence alignments, and for the
challenging new DockGen-E dataset, FlowDock outperforms single-sequence
AlphaFold 3 and matches single-sequence Chai-1 for binding pocket
generalization. Additionally, in the ligand category of the 16th community-wide
Critical Assessment of Techniques for Structure Prediction (CASP16), FlowDock
ranked among the top-5 methods for pharmacological binding affinity estimation
across 140 protein-ligand complexes, demonstrating the efficacy of its learned
representations in virtual screening. Source code, data, and pre-trained models
are available at https://github.com/BioinfoMachineLearning/FlowDock.


## AI 摘要

FlowDock是一种基于条件流匹配的深度几何生成模型，首次实现了将未结合（apo）蛋白质结构直接映射到任意数量配体结合的（holo）结构。它能同时预测多个配体的结合构象，并提供结构置信度和结合亲和力评分，支持快速虚拟筛选。在PoseBusters基准测试中，FlowDock以51%的盲对接成功率超越单序列AlphaFold 3；在DockGen-E数据集上，其结合口袋泛化能力与Chai-1相当。在CASP16竞赛中，FlowDock位列前5名，展现了其在药物发现中的潜力。代码、数据和预训练模型已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T15:02:29Z
- **目录日期**: 2025-03-25

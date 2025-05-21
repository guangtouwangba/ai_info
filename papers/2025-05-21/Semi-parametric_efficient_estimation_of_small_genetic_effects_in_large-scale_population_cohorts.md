# Semi-parametric efficient estimation of small genetic effects in large-scale population cohorts

**URL**: http://arxiv.org/abs/2505.14675v1

## 原始摘要

Population genetics seeks to quantify DNA variant associations with traits or
diseases, as well as interactions among variants and with environmental
factors. Computing millions of estimates in large cohorts in which small effect
sizes are expected, necessitates minimising model-misspecification bias to
control false discoveries. We present TarGene, a unified statistical workflow
for the semi-parametric efficient and double robust estimation of genetic
effects including k-point interactions among categorical variables in the
presence of confounding and weak population dependence. k-point interactions,
or Average Interaction Effects (AIEs), are a direct generalisation of the usual
average treatment effect (ATE). We estimate AIEs with cross-validated and/or
weighted versions of Targeted Minimum Loss-based Estimators (TMLE) and One-Step
Estimators (OSE). The effect of dependence among data units on variance
estimates is corrected by using sieve plateau variance estimators based on
genetic relatedness across the units. We present extensive realistic
simulations to demonstrate power, coverage, and control of type I error. Our
motivating application is the targeted estimation of genetic effects on trait,
including two-point and higher-order gene-gene and gene-environment
interactions, in large-scale genomic databases such as UK Biobank and All of
Us. All cross-validated and/or weighted TMLE and OSE for the AIE k-point
interaction, as well as ATEs, conditional ATEs and functions thereof, are
implemented in the general purpose Julia package TMLE.jl. For high-throughput
applications in population genomics, we provide the open-source Nextflow
pipeline and software TarGene which integrates seamlessly with modern
high-performance and cloud computing platforms.


## AI 摘要

TarGene是一个统一的统计工作流程，用于半参数高效且双重稳健地估计遗传效应，包括存在混杂因素和弱群体依赖性时分类变量的k点交互作用（平均交互效应，AIE）。它采用交叉验证/加权的目标最小损失估计器（TMLE）和一步估计器（OSE），并通过基于遗传相关性的筛板方差估计器校正数据依赖性对方差的影响。该方法适用于大规模基因组数据库（如UK Biobank），能精准估计基因-基因/基因-环境交互作用，控制I类错误并保持统计功效。相关算法已集成至Julia包TMLE.jl和开源流程TarGene，支持高性能与云计算平台。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T10:01:17Z
- **目录日期**: 2025-05-21

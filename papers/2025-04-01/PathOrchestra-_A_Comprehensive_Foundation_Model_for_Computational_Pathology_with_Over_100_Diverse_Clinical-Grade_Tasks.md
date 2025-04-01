# PathOrchestra: A Comprehensive Foundation Model for Computational Pathology with Over 100 Diverse Clinical-Grade Tasks

**URL**: http://arxiv.org/abs/2503.24345v1

## 原始摘要

The complexity and variability inherent in high-resolution pathological
images present significant challenges in computational pathology. While
pathology foundation models leveraging AI have catalyzed transformative
advancements, their development demands large-scale datasets, considerable
storage capacity, and substantial computational resources. Furthermore,
ensuring their clinical applicability and generalizability requires rigorous
validation across a broad spectrum of clinical tasks. Here, we present
PathOrchestra, a versatile pathology foundation model trained via
self-supervised learning on a dataset comprising 300K pathological slides from
20 tissue and organ types across multiple centers. The model was rigorously
evaluated on 112 clinical tasks using a combination of 61 private and 51 public
datasets. These tasks encompass digital slide preprocessing, pan-cancer
classification, lesion identification, multi-cancer subtype classification,
biomarker assessment, gene expression prediction, and the generation of
structured reports. PathOrchestra demonstrated exceptional performance across
27,755 WSIs and 9,415,729 ROIs, achieving over 0.950 accuracy in 47 tasks,
including pan-cancer classification across various organs, lymphoma subtype
diagnosis, and bladder cancer screening. Notably, it is the first model to
generate structured reports for high-incidence colorectal cancer and
diagnostically complex lymphoma-areas that are infrequently addressed by
foundational models but hold immense clinical potential. Overall, PathOrchestra
exemplifies the feasibility and efficacy of a large-scale, self-supervised
pathology foundation model, validated across a broad range of clinical-grade
tasks. Its high accuracy and reduced reliance on extensive data annotation
underline its potential for clinical integration, offering a pathway toward
more efficient and high-quality medical services.


## AI 摘要

PathOrchestra是一个基于自监督学习的大规模病理基础模型，利用30万张多中心病理切片训练而成。该模型在112项临床任务（包括癌症分类、病变识别、生物标志物评估等）中表现优异，在27,755张全切片图像和940万区域上验证，47项任务准确率超过0.950。其创新性在于首次实现了高发结直肠癌和复杂淋巴瘤的结构化报告生成。研究表明该模型具有临床级准确性和泛化能力，显著降低对数据标注的依赖，为高效、高质量的医疗服务提供了可行路径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T18:02:14Z
- **目录日期**: 2025-04-01

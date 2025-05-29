# Chain-of-Talkers (CoTalk): Fast Human Annotation of Dense Image Captions

**URL**: http://arxiv.org/abs/2505.22627v1

## 原始摘要

While densely annotated image captions significantly facilitate the learning
of robust vision-language alignment, methodologies for systematically
optimizing human annotation efforts remain underexplored. We introduce
Chain-of-Talkers (CoTalk), an AI-in-the-loop methodology designed to maximize
the number of annotated samples and improve their comprehensiveness under fixed
budget constraints (e.g., total human annotation time). The framework is built
upon two key insights. First, sequential annotation reduces redundant workload
compared to conventional parallel annotation, as subsequent annotators only
need to annotate the ``residual'' -- the missing visual information that
previous annotations have not covered. Second, humans process textual input
faster by reading while outputting annotations with much higher throughput via
talking; thus a multimodal interface enables optimized efficiency. We evaluate
our framework from two aspects: intrinsic evaluations that assess the
comprehensiveness of semantic units, obtained by parsing detailed captions into
object-attribute trees and analyzing their effective connections; extrinsic
evaluation measures the practical usage of the annotated captions in
facilitating vision-language alignment. Experiments with eight participants
show our Chain-of-Talkers (CoTalk) improves annotation speed (0.42 vs. 0.30
units/sec) and retrieval performance (41.13\% vs. 40.52\%) over the parallel
method.


## AI 摘要

本文提出Chain-of-Talkers (CoTalk)方法，通过AI辅助优化图像标注效率。该方法采用顺序标注策略，后续标注者只需补充前序遗漏信息，减少冗余工作；同时结合多模态界面（阅读输入+语音输出）提升效率。实验表明，在固定预算下，CoTalk相比并行标注速度更快（0.42 vs. 0.30单位/秒），生成的标注更全面，且能提升视觉-语言对齐任务的检索性能（41.13% vs. 40.52%）。内在评估通过解析对象-属性树分析语义完整性，外在评估验证了标注对下游任务的实际促进效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T03:21:10Z
- **目录日期**: 2025-05-29

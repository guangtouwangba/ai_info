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

本文提出Chain-of-Talkers (CoTalk)方法，在固定预算下优化图像标注效率。该方法采用AI辅助的串行标注模式，后续标注者只需补充前序遗漏的视觉信息，减少冗余工作。同时结合多模态界面（阅读输入+语音输出）提升效率。实验表明：CoTalk的标注速度（0.42单位/秒）优于传统并行方法（0.30单位/秒），其生成的标注在语义完整性（通过对象-属性树分析）和实际应用（视觉-语言对齐任务中检索准确率达41.13%）方面均表现更优。该方法通过8人实验验证了有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T11:02:03Z
- **目录日期**: 2025-05-29

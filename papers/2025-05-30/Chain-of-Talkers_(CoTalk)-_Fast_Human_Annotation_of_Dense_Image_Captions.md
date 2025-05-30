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

CoTalk提出了一种AI辅助的序列标注方法，在固定预算下最大化标注样本数量和全面性。其核心创新在于：1）通过序列标注减少冗余工作量，后序标注者只需补充前序遗漏的视觉信息；2）采用多模态界面（阅读输入+语音输出）提升效率。实验表明，相比并行标注，CoTalk的标注速度提升40%（0.42单位/秒），且通过对象-属性树分析证实标注更全面，在视觉-语言对齐任务中检索准确率提高0.61%（41.13%）。该方法通过8人实验验证了优化人机协作标注流程的有效性。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T00:02:35Z
- **目录日期**: 2025-05-30

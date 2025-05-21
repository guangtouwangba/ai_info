# NExT-Search: Rebuilding User Feedback Ecosystem for Generative AI Search

**URL**: http://arxiv.org/abs/2505.14680v1

## 原始摘要

Generative AI search is reshaping information retrieval by offering
end-to-end answers to complex queries, reducing users' reliance on manually
browsing and summarizing multiple web pages. However, while this paradigm
enhances convenience, it disrupts the feedback-driven improvement loop that has
historically powered the evolution of traditional Web search. Web search can
continuously improve their ranking models by collecting large-scale,
fine-grained user feedback (e.g., clicks, dwell time) at the document level. In
contrast, generative AI search operates through a much longer search pipeline,
spanning query decomposition, document retrieval, and answer generation, yet
typically receives only coarse-grained feedback on the final answer. This
introduces a feedback loop disconnect, where user feedback for the final output
cannot be effectively mapped back to specific system components, making it
difficult to improve each intermediate stage and sustain the feedback loop. In
this paper, we envision NExT-Search, a next-generation paradigm designed to
reintroduce fine-grained, process-level feedback into generative AI search.
NExT-Search integrates two complementary modes: User Debug Mode, which allows
engaged users to intervene at key stages; and Shadow User Mode, where a
personalized user agent simulates user preferences and provides AI-assisted
feedback for less interactive users. Furthermore, we envision how these
feedback signals can be leveraged through online adaptation, which refines
current search outputs in real-time, and offline update, which aggregates
interaction logs to periodically fine-tune query decomposition, retrieval, and
generation models. By restoring human control over key stages of the generative
AI search pipeline, we believe NExT-Search offers a promising direction for
building feedback-rich AI search systems that can evolve continuously alongside
human feedback.


## AI 摘要

生成式AI搜索通过提供端到端答案简化了信息检索，但也破坏了传统网络搜索依赖用户反馈（如点击、停留时间）持续改进的循环机制。由于AI搜索流程复杂（包括查询分解、文档检索和答案生成），而用户仅能对最终结果提供粗粒度反馈，导致难以优化中间环节。为此，研究提出NExT-Search新范式，通过两种模式重新引入细粒度反馈：用户调试模式（允许干预关键阶段）和影子用户模式（AI代理模拟用户偏好）。该系统支持实时在线调整和离线模型更新，旨在重建人机协同的持续改进闭环。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T15:01:06Z
- **目录日期**: 2025-05-21

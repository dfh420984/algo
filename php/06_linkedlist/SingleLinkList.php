<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/21
 * Time: 14:56
 */

/**
 * 单向链表
 */

require_once "./SingleLinkListNode.php";

class SingleLinkList {
    public $head; //哨兵头节点
    public $length;

    public function __construct($head = null)
    {
        if ($head == null) {
            $this->head = new SingleLinkListNode();
        } else {
            $this->head = $head;
        }
        $this->length = 0;
    }

    //顺序插入节点
    public function insert($data) {
        return $this->insertDataAfterNode($this->head,$data);
    }

    //在某个节点后面插入新节点数据（指定节点插入）
    public function insertDataAfterNode($originNode,$data) {
        $newNode = new SingleLinkListNode($data); //新节点
        $newNode->data = $data;
        $newNode->next = $originNode->next;
        if ($originNode == null) {
            return false;
        } else  {
            $originNode->next = $newNode;
        }
        $this->length++;
        return $newNode;
    }

    //根据值查找节点
    public function searchNode($val) {
        $cur = $this->head;
        while ($cur != null) {
            if ($cur->data ==  $val) {
                return $cur;
            } else {
                $cur = $cur->next;
            }
        }
        return null;
    }

    //获取当前节点的前一个节点
    public function getPreNode($node) {
        if (null == $node) {
            return false;
        }
        $cur = $this->head;
        $pre = $this->head;
        while ($cur != null) {
            if ($cur->next === $node) {
                $pre =  $cur;
                return $pre;
            } else {
                $cur = $cur->next;
            }
        }
        return $pre;
    }

    /**
     * 通过索引获取节点
     *
     * @param int $index
     *
     * @return SingleLinkedListNode|null
     */
    public function getNodeByIndex($index)
    {
        if ($index >= $this->length) {
            return null;
        }
        $cur = $this->head->next;
        for ($i = 0; $i < $index; ++$i) {
            $cur = $cur->next;
        }
        return $cur;
    }

    //删除指定数据节点
    public function deleteNode($val) {
        $cur = $this->searchNode($val);
        if ($cur == null) { //找不到节点
            return false;
        } else {
            //获取前一个节点
            $pre = $this->getPreNode($cur);
            $pre->next = $cur->next;
            $cur = null;
            $this->length--;
            return true;
        }
    }

    //输出联表节点信息
    public function printNode() {
        if ($this->head->next == null) {
            return false;
        }
        $cur = $this->head;
        while ($cur->next != null) {
            echo $cur->next->data . ' -> ';
            $cur = $cur->next;
        }
        echo "NULL".PHP_EOL;
        return true;
    }

    //输出带环节点信息
    public function printCircleNode() {
        if ($this->head->next == null) {
            return false;
        }
        $cur = $this->head;
        $length = $this->getLength();
        while ($cur->next != null && $length--) {
            echo $cur->next->data . ' -> ';
            $cur = $cur->next;
        }
        echo "NULL".PHP_EOL;
        return true;
    }

    /**
     * 在某个节点后插入新的节点
     *
     * @param SingleLinkedListNode $originNode
     * @param SingleLinkedListNode $node
     *
     * @return SingleLinkedListNode|bool
     */
    public function insertNodeAfter($originNode, $node)
    {
        // 如果originNode为空，插入失败
        if (null == $originNode) {
            return false;
        }
        $node->next = $originNode->next;
        $originNode->next = $node;
        $this->length++;
        return $node;
    }

    //获取联表长度
    public function getLength() {
        return $this->length;
    }

    /**
     * 定义两个指针，同时从链表的头节点出发，一个指针一次走一步，另一个指针一次走两步。如果走得快的指针追上了走得慢的指针，
     * 那么链表就是环形链表；如果走得快的指针走到了链表的末尾（next指向 NULL）都没有追上第一个指针，那么链表就不是环形链表
     * @param $head 带头节点的单链表
     * @author duanfuhao@smzdm.com
     * @date 2019/8/21
     */
    public function isCircleLink($head) {
        if ($head == null) {
            return false;
        }
        $slow = $head->next; // 初始时，慢指针从头节点开始走1步
        if ($slow == null) {
            return false;
        }
        $fast = $slow->next; // 初始时，快指针从头节点开始走2步
        while ($fast != null && $slow != null ) { //如果没环，走到链尾就结束循环
            if ($fast == $slow) { //快慢指针相遇，说明有环
                return true;
            }
            $slow = $slow->next;
            $fast = $fast->next;
            if ($fast != null) {
                $fast = $fast->next;
            }
        }
        return false;
    }

    /**
     * 构造一个有环的链表
     */
    public function buildHasCircleList()
    {
        $data = [1, 2, 3, 4, 5, 6, 7, 8];
        $node0 = new SingleLinkListNode($data[0]);
        $node1 = new SingleLinkListNode($data[1]);
        $node2 = new SingleLinkListNode($data[2]);
        $node3 = new SingleLinkListNode($data[3]);
        $node4 = new SingleLinkListNode($data[4]);
        $node5 = new SingleLinkListNode($data[5]);
        $node6 = new SingleLinkListNode($data[6]);
        $node7 = new SingleLinkListNode($data[7]);
        $this->insertNodeAfter($this->head, $node0);
        $this->insertNodeAfter($node0, $node1);
        $this->insertNodeAfter($node1, $node2);
        $this->insertNodeAfter($node2, $node3);
        $this->insertNodeAfter($node3, $node4);
        $this->insertNodeAfter($node4, $node5);
        $this->insertNodeAfter($node5, $node6);
        $this->insertNodeAfter($node6, $node7);
        $node7->next = $node0;
    }

    /**
     * @param $head 第一个节点
     * @author duanfuhao@smzdm.com
     * @date 2019/8/22
     */
    function reverseLinkedList($head) {
        $pre = null; //指向前一个节点
        $next = null; //下一个节点
        while ($head != null) {
            $next = $head->next; //下一个节点
            $head->next = $pre;  //修改节点指向位反转后的前节点
            $pre = $head; //将pre指向当前方转后的节点
            $head = $next; //head指针后移，开始下一轮循环反转
        }
        return $pre;
    }

    //判断链表中保存的字符串是否是回文
    function isLinkPalindrome($list) {
        if ($list->getLength() <= 1) {
            return false;
        }
        //定义快慢两个指针
        $slow = $list->head->next;
        $fast = $list->head->next;

        $midNode = null; //中间点
        $pre = null; //反转后链表对像

        // 找单链表中点
        while ($fast != null && $fast->next != null  && $fast->next->next != null) {
            $fast = $fast->next->next;
            $midNode = $slow->next;
        }
        //如果是偶数
        if ($this->getLength() % 2 == 0) {
            $midNode = $midNode->next;
        }
        //反转后半部分链表
        while ($midNode != null) {
            $next = $midNode->next;
            $midNode->next = $pre;
            $pre = $midNode;
            $midNode = $next;
        }
        // 开始逐个比较
        while ($slow != null) {
            if ($slow->data != $pre->data) {
                return false;
            }
            $slow = $slow->next;
            $pre = $pre->next;
        }
        return true;
    }
}
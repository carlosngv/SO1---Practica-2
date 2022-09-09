#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>
/* Header para usar la lib seq_file y manejar el w_file en /proc*/
#include <linux/seq_file.h>

/* libreria de memoria ram*/
#include <linux/hugetlb.h>

#include <linux/sched.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("CPU module");
MODULE_AUTHOR("Carlos Ng");

struct task_struct * cpu;
struct task_struct * child;
struct list_head * lstProcess;


static int write_file(struct seq_file *w_file, void *v)
{

/* Structure

    {
        "cpu": [
            {
                "pid": pid,
                "comm": comm,
                "children": [
                    {
                        "childPid": childPid,
                        "childComm": childComm
                    }
                ]
            }
        ]
    }

*/
    seq_printf(w_file, "{\n");
    seq_printf(w_file, "   ");
    seq_printf(w_file, "\"cpu\": [\n");
    for_each_process(cpu){
        seq_printf(w_file, "       ");
        seq_printf(w_file, "{\n");
        seq_printf(w_file, "           ");
        seq_printf(w_file, "\"pid\": ");
        seq_printf(w_file, "\"%d\",\n", cpu->pid);
        seq_printf(w_file, "           ");
        seq_printf(w_file, "\"comm\": ");
        seq_printf(w_file, "\"%s\",\n", cpu->comm);
        seq_printf(w_file, "           ");
        seq_printf(w_file, "\"children\": [\n");
        list_for_each(lstProcess, &(cpu->children)){
            // Procesos hijos
            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(w_file, "                   ");
            seq_printf(w_file, "{\n");
            seq_printf(w_file, "                       ");
            seq_printf(w_file, "\"pid\": ");
            seq_printf(w_file, "\"%d\",\n", child->pid);
            seq_printf(w_file, "                       ");
            seq_printf(w_file, "\"comm\": ");
            seq_printf(w_file, "\"%s\",\n", child->comm);
            seq_printf(w_file, "\n");
            seq_printf(w_file, "                   ");
            seq_printf(w_file, "},\n");
        }
        seq_printf(w_file, "           ");
        seq_printf(w_file, "]\n");
        seq_printf(w_file, "       ");
        seq_printf(w_file, "},\n");
    }
    seq_printf(w_file, "   ");
    seq_printf(w_file, "]\n");
    seq_printf(w_file, "}\n");
    return 0;
}


static int open(struct inode *inode, struct file *file)
{
    return single_open(file, write_file, NULL);
}


static struct proc_ops operations =
{
    .proc_open = open,
    .proc_read = seq_read
};


static int _insert(void)
{

    // RAM_carnet o CPU_carnet
    proc_create("ram_201801434", 0, NULL, &operations);
    printk(KERN_INFO "201801434\n");
    return 0;
}


static void _remove(void)
{
    remove_proc_entry("ram_201801434", NULL);
    printk(KERN_INFO "Sistemas Operativos 1\n");
}

module_init(_insert);
module_exit(_remove);
